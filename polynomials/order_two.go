package polynomials

import (
	"math/big"

	"github.com/gbdubs/maybebig/maybebig"
	"github.com/gbdubs/polynomials/bigcomplex"
	"github.com/gbdubs/polynomials/maybebigcomplex"
)

func SecondOrder(a, b, c *bigcomplex.BigComplex) []*bigcomplex.BigComplex {
	if a.IsZero() {
		return FirstOrder(b, c)
	}
	disc := sub(sq(b), mul(mul(a, c), bigcomplex.FromInt(4)))
	negOne := bigcomplex.FromInt(-1)
	two := bigcomplex.FromInt(2)
	if disc.IsZero() {
		return []*bigcomplex.BigComplex{div(mul(negOne, b), mul(two, a))}
	}
	root := disc.Sqrt()
	result := []*bigcomplex.BigComplex{
		div(sub(mul(negOne, b), root), mul(two, a)),
		div(add(mul(negOne, b), root), mul(two, a)),
	}
	return bigcomplex.Unique(result)
}

func SecondOrderMB(a, b, c *maybebigcomplex.BigComplex) []*maybebigcomplex.BigComplex {
	if a.IsZero() {
		return FirstOrderMB(b, c)
	}
	disc := subMB(sqMB(b), mulMB(mulMB(a, c), maybebigcomplex.FromInt(4)))
	negOne := maybebigcomplex.FromInt(-1)
	two := maybebigcomplex.FromInt(2)
	if disc.IsZero() {
		return []*maybebigcomplex.BigComplex{divMB(mulMB(negOne, b), mulMB(two, a))}
	}
	root := disc.Sqrt()
	result := []*maybebigcomplex.BigComplex{
		divMB(subMB(mulMB(negOne, b), root), mulMB(two, a)),
		divMB(addMB(mulMB(negOne, b), root), mulMB(two, a)),
	}
	return maybebigcomplex.Unique(result)
}

func SecondOrderComplex128(a, b, c complex128) []complex128 {
	return ToComplex128s(
		SecondOrder(
			FromComplex128UseWithCaution(a),
			FromComplex128UseWithCaution(b),
			FromComplex128UseWithCaution(c)))
}

func SecondOrderReal(a, b, c *big.Float) []*big.Float {
	return RealComponents(
		FilterToReals(
			SecondOrder(
				&bigcomplex.BigComplex{Real: newFloat().Set(a), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(b), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(c), Imag: newFloat()})))
}

func SecondOrderRealMB(a, b, c *maybebig.Float) []*maybebig.Float {
	return RealComponentsMB(
		FilterToRealsMB(
			SecondOrderMB(
				&maybebigcomplex.BigComplex{Real: a, Imag: maybebig.NewFloatZ()},
				&maybebigcomplex.BigComplex{Real: b, Imag: maybebig.NewFloatZ()},
				&maybebigcomplex.BigComplex{Real: c, Imag: maybebig.NewFloatZ()})))
}

func SecondOrderRealFloat64(a, b, c float64) []float64 {
	return Float64s(
		RealComponents(
			FilterToReals(
				SecondOrder(
					FromFloat64UseWithCaution(a),
					FromFloat64UseWithCaution(b),
					FromFloat64UseWithCaution(c)))))
}
