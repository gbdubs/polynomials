package bigcomplex

import (
	"math/big"

	"github.com/gbdubs/polynomials/precision"
)

func newFloat() *big.Float {
	f := big.NewFloat(0)
	f.SetPrec(precision.Get())
	f.SetMode(big.ToZero)
	return f
}

func add(fs ...*big.Float) *big.Float {
	r := newFloat()
	for _, f := range fs {
		r.Add(r, f)
	}
	return r
}

func mul(fs ...*big.Float) *big.Float {
	result := newFloat().SetInt64(1)
	for _, f := range fs {
		result = product(result, f)
	}
	return result
}

func product(a, b *big.Float) *big.Float {
	return newFloat().Mul(a, b)
}

func sub(a, b *big.Float) *big.Float {
	return newFloat().Sub(a, b)
}

func div(a, b *big.Float) *big.Float {
	return newFloat().Quo(a, b)
}

func sq(a *big.Float) *big.Float {
	return mul(a, a)
}

func sqrtSimple(a *big.Float) *big.Float {
	if isZero(a) {
		return newFloat()
	}
	return newFloat().Sqrt(a)
}

func abs(a *big.Float) *big.Float {
	return newFloat().Abs(a)
}

func zeroForCmp() *big.Float {
	return big.NewFloat(0).SetPrec(precision.Get() / 2).SetMode(big.ToZero)
}

func isZero(a *big.Float) bool {
	a = zeroForCmp().Set(a)
	abs := abs(a)
	return zeroForCmp().Cmp(a) == 0 || abs.Cmp(big.NewFloat(1e-300)) < 0
}

func positive(a *big.Float) bool {
	a = zeroForCmp().Set(a)
	return zeroForCmp().Cmp(a) < 0
}

func negative(a *big.Float) bool {
	a = zeroForCmp().Set(a)
	return zeroForCmp().Cmp(a) > 0
}

func fromInt(i int) *big.Float {
	return newFloat().SetInt64(int64(i))
}

func FromInt(i int) *BigComplex {
	return &BigComplex{
		Real: newFloat().SetInt64(int64(i)),
		Imag: newFloat(),
	}
}
