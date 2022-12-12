package bigtrig

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
	return newFloat().Sqrt(a)
}

func abs(a *big.Float) *big.Float {
	return newFloat().Abs(a)
}

func isZero(a *big.Float) bool {
	return newFloat().Cmp(a) == 0
}

func positive(a *big.Float) bool {
	return newFloat().Cmp(a) < 0
}

func negative(a *big.Float) bool {
	return newFloat().Cmp(a) > 0
}
