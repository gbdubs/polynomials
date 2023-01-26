package polynomials

import (
	"math/big"

	"github.com/gbdubs/polynomials/bigcomplex"
	"github.com/gbdubs/polynomials/maybebigcomplex"
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

func addMB(bcs ...*maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return maybebigcomplex.Add(bcs...)
}
func subMB(a, b *maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return maybebigcomplex.Sub(a, b)
}
func mulMB(bcs ...*maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return maybebigcomplex.Mul(bcs...)
}
func divMB(a, b *maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return maybebigcomplex.Div(a, b)
}
func sqMB(a *maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return maybebigcomplex.Sq(a)
}
func cubeMB(a *maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return maybebigcomplex.Cube(a)
}
func sqrtMB(a *maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return a.Sqrt()
}
func cbrtMB(a *maybebigcomplex.BigComplex) *maybebigcomplex.BigComplex {
	return a.Cbrt()
}
func fromIntMB(i int) *maybebigcomplex.BigComplex {
	return maybebigcomplex.FromInt(i)
}
func newFloatmB() *big.Float {
	return big.NewFloat(0).SetPrec(precision.Get())
}
func newRatMB(a, b int) *big.Float {
	return newFloat().SetRat(big.NewRat(int64(a), int64(b)))
}
