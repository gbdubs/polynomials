package ivyshims

import (
	"fmt"
	"math/big"

	"github.com/gbdubs/polynomials/precision"
)

func Atan(q *big.Float) *big.Float {
	return bigFloat(atan(shimContext(), shimValue(q)))
}

func Cos(q *big.Float) *big.Float {
	return bigFloat(cos(shimContext(), shimValue(q)))
}

func Sin(q *big.Float) *big.Float {
	return bigFloat(sin(shimContext(), shimValue(q)))
}

func Pow(x *big.Float, e *big.Float) *big.Float {
	return bigFloat(power(shimContext(), shimValue(x), shimValue(e)))
}

func shimContext() Context {
	c := &Config{}
	c.init()
	c.SetFloatPrec(precision.Get())
	return NewContext(c)
}

func shimValue(b *big.Float) Value {
	return BigFloat{b}
}

func bigFloat(v Value) *big.Float {
	if u, ok := v.(BigFloat); ok {
		return u.Float
	}
	panic(fmt.Errorf("expected to return big float but did not"))
}
