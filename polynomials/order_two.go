package polynomials

import (
	"math/big"

	"github.com/gbdubs/polynomials/bigcomplex"
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
	roots := disc.Sqrt()
	result := []*bigcomplex.BigComplex{}
	for _, root := range roots {
		result = append(result, div(sub(mul(negOne, b), root), mul(two, a)))
		result = append(result, div(add(mul(negOne, b), root), mul(two, a)))
	}
	return bigcomplex.Unique(result)
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

func SecondOrderRealFloat64(a, b, c float64) []float64 {
	return Float64s(
		RealComponents(
			FilterToReals(
				SecondOrder(
					FromFloat64UseWithCaution(a),
					FromFloat64UseWithCaution(b),
					FromFloat64UseWithCaution(c)))))
}
