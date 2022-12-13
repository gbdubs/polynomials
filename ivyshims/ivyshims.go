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

func ComplexSqrt(r, i *big.Float) (*big.Float, *big.Float) {
	v := Complex{real: shimValue(r), imag: shimValue(i)}
	res := complexSqrt(shimContext(), v)
	return bigFloat(res.real), bigFloat(res.imag)
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
	if u, ok := v.(Int); ok {
		return newFloat(shimContext()).SetInt64(int64(u))
	}
	panic(fmt.Errorf("expected to return BigFloat but returned %T: %+v", v, v))
}
