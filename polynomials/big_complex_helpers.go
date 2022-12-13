package polynomials

import (
	"math/big"

	"github.com/gbdubs/polynomials/bigcomplex"
	"github.com/gbdubs/polynomials/precision"
)

func add(bcs ...*bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return bigcomplex.Add(bcs...)
}
func sub(a, b *bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return bigcomplex.Sub(a, b)
}
func mul(bcs ...*bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return bigcomplex.Mul(bcs...)
}
func div(a, b *bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return bigcomplex.Div(a, b)
}
func sq(a *bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return bigcomplex.Sq(a)
}
func cube(a *bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return bigcomplex.Cube(a)
}
func sqrt(a *bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return a.Sqrt()
}
func cbrt(a *bigcomplex.BigComplex) *bigcomplex.BigComplex {
	return a.Cbrt()
}
func fromInt(i int) *bigcomplex.BigComplex {
	return bigcomplex.FromInt(i)
}
func newFloat() *big.Float {
	return big.NewFloat(0).SetPrec(precision.Get())
}
func newRat(a, b int) *big.Float {
	return newFloat().SetRat(big.NewRat(int64(a), int64(b)))
}
