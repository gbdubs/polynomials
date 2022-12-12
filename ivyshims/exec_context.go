// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ivyshims // import "robpike.io/ivy/exec"

import (
	"strings"
)

// Symtab is a symbol table, a map of names to values.
type Symtab map[string]Value

// Context holds execution context, specifically the binding of names to values and operators.
// It is the only implementation of ../Context, but since it references the value
// package, there would be a cycle if that package depended on this type definition.
type ExecContext struct {
	// config is the configuration state used for evaluation, printing, etc.
	// Accessed through the Context Config method.
	config *Config

	frameSizes []int // size of each stack frame on the call stack
	stack      []Value

	Globals Symtab

	//  UnaryFn maps the names of unary functions (ops) to their implemenations.
	UnaryFn map[string]*Function
	//  BinaryFn maps the names of binary functions (ops) to their implemenations.
	BinaryFn map[string]*Function
	// Defs is a list of defined ops, in time order.  It is used when saving the
	// Context to a file.
	Defs []OpDef
	// Names of variables declared in the currently-being-parsed function.
	variables []string
}

// NewContext returns a new execution context: the stack and variables,
// plus the execution configuration.
func NewContext(conf *Config) Context {
	c := &ExecContext{
		config:   conf,
		Globals:  make(Symtab),
		UnaryFn:  make(map[string]*Function),
		BinaryFn: make(map[string]*Function),
	}
	c.SetConstants()
	return c
}

func (c *ExecContext) Config() *Config {
	return c.config
}

// SetConstants re-assigns the fundamental constant values using the current
// setting of floating-point precision.
func (c *ExecContext) SetConstants() {
	e, pi := Consts(c)
	c.AssignGlobal("e", e)
	c.AssignGlobal("pi", pi)
}

// Global returns the value of a global symbol, or nil if the symbol is not defined globally.
func (c *ExecContext) Global(name string) Value {
	return c.Globals[name]
}

// Local returns the value of the local variable with index i.
func (c *ExecContext) Local(i int) Value {
	return c.stack[len(c.stack)-i]
}

// AssignLocal assigns the local variable with the given index the value.
func (c *ExecContext) AssignLocal(i int, value Value) {
	c.stack[len(c.stack)-i] = value
}

// Assign assigns the global variable the value. The variable must
// be defined either in the current function or globally.
// Inside a function, new variables become locals.
func (c *ExecContext) AssignGlobal(name string, val Value) {
	c.Globals[name] = val
}

// push pushes a new local frame onto the context stack.
func (c *ExecContext) push(fn *Function) {
	n := len(c.stack)
	for cap(c.stack) < n+len(fn.Locals) {
		c.stack = append(c.stack[:cap(c.stack)], nil)
	}
	c.frameSizes = append(c.frameSizes, len(fn.Locals))
	c.stack = c.stack[:n+len(fn.Locals)]
}

// pop pops the top frame from the stack.
func (c *ExecContext) pop() {
	n := c.frameSizes[len(c.frameSizes)-1]
	c.frameSizes = c.frameSizes[:len(c.frameSizes)-1]
	c.stack = c.stack[:len(c.stack)-n]
}

// Eval evaluates a list of expressions.
func (c *ExecContext) Eval(exprs []Expr) []Value {
	var values []Value
	for _, expr := range exprs {
		v := expr.Eval(c)
		if v != nil {
			values = append(values, v)
		}
	}
	return values
}

// EvalUnary evaluates a unary operator, including reductions and scans.
func (c *ExecContext) EvalUnary(op string, right Value) Value {
	if len(op) > 1 {
		switch op[len(op)-1] {
		case '/':
			return Reduce(c, op[:len(op)-1], right)
		case '\\':
			return Scan(c, op[:len(op)-1], right)
		}
	}
	fn := c.Unary(op)
	if fn == nil {
		Errorf("unary %q not implemented", op)
	}
	return fn.EvalUnary(c, right)
}

func (c *ExecContext) Unary(op string) UnaryOp {
	userFn := c.UnaryFn[op]
	if userFn != nil {
		return userFn
	}
	builtin := UnaryOps[op]
	if builtin != nil {
		return builtin
	}
	return nil
}

func (c *ExecContext) UserDefined(op string, isBinary bool) bool {
	if isBinary {
		return c.BinaryFn[op] != nil
	}
	return c.UnaryFn[op] != nil
}

// EvalBinary evaluates a binary operator, including products.
func (c *ExecContext) EvalBinary(left Value, op string, right Value) Value {
	if strings.Contains(op, ".") {
		return Product(c, left, op, right)
	}
	fn := c.Binary(op)
	if fn == nil {
		Errorf("binary %q not implemented", op)
	}
	return fn.EvalBinary(c, left, right)
}

func (c *ExecContext) Binary(op string) BinaryOp {
	user := c.BinaryFn[op]
	if user != nil {
		return user
	}
	builtin := BinaryOps[op]
	if builtin != nil {
		return builtin
	}
	return nil
}

// Define defines the function and installs it. It also performs
// some error checking and adds the function to the sequencing
// information used by the save method.
func (c *ExecContext) Define(fn *Function) {
	c.noVar(fn.Name)
	if fn.IsBinary {
		c.BinaryFn[fn.Name] = fn
	} else {
		c.UnaryFn[fn.Name] = fn
	}
	// Update the sequence of definitions.
	// First, if it's last (a very common case) there's nothing to do.
	if len(c.Defs) > 0 {
		last := c.Defs[len(c.Defs)-1]
		if last.Name == fn.Name && last.IsBinary == fn.IsBinary {
			return
		}
	}
	// Is it already defined?
	for i, def := range c.Defs {
		if def.Name == fn.Name && def.IsBinary == fn.IsBinary {
			// Yes. Drop it.
			c.Defs = append(c.Defs[:i], c.Defs[i+1:]...)
			break
		}
	}
	// It is now the most recent definition.
	c.Defs = append(c.Defs, OpDef{fn.Name, fn.IsBinary})
}

// noVar guarantees that there is no global variable with that name,
// preventing an op from being defined with the same name as a variable,
// which could cause problems. A variable with value zero is considered to
// be OK, so one can clear a variable before defining a symbol. A cleared
// variable is removed from the global symbol table.
// noVar also prevents defining builtin variables as ops.
func (c *ExecContext) noVar(name string) {
	if name == "_" || name == "pi" || name == "e" { // Cannot redefine these.
		Errorf(`cannot define op with name %q`, name)
	}
	sym := c.Globals[name]
	if sym == nil {
		return
	}
	if i, ok := sym.(Int); ok && i == 0 {
		delete(c.Globals, name)
		return
	}
	Errorf("cannot define op %s; it is a variable (%[1]s=0 to clear)", name)
}

// noOp is the dual of noVar. It also checks for assignment to builtins.
// It just errors out if there is a conflict.
func (c *ExecContext) noOp(name string) {
	if name == "pi" || name == "e" { // Cannot redefine these.
		Errorf("cannot reassign %q", name)
	}
	if c.UnaryFn[name] == nil && c.BinaryFn[name] == nil {
		return
	}
	Errorf("cannot define variable %s; it is an op", name)
}

// Declare makes the name a variable while parsing the next function.
func (c *ExecContext) Declare(name string) {
	c.variables = append(c.variables, name)
}

// ForgetAll forgets the declared variables.
func (c *ExecContext) ForgetAll() {
	c.variables = nil
}

func (c *ExecContext) isVariable(op string) bool {
	for _, s := range c.variables {
		if op == s {
			return true
		}
	}
	return false
}
