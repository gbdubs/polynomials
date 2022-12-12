package polynomials

import (
	"math/big"

	"github.com/gbdubs/polynomials/bigcomplex"
)

func FirstOrder(a, b *bigcomplex.BigComplex) []*bigcomplex.BigComplex {
	if a.IsZero() {
		// No solution possible, either allways satisfies or nothing satisfies
		return nil
	}
	return []*bigcomplex.BigComplex{
		div(mul(bigcomplex.FromInt(-1), b), a),
	}
}

func FirstOrderComplex128(a, b complex128) []complex128 {
	return ToComplex128s(
		FirstOrder(
			FromComplex128UseWithCaution(a),
			FromComplex128UseWithCaution(b)))
}

func FirstOrderReal(a, b, c, d, e *big.Float) []*big.Float {
	return RealComponents(
		FilterToReals(
			FirstOrder(
				&bigcomplex.BigComplex{Real: newFloat().Set(a), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(b), Imag: newFloat()})))
}

func FirstOrderRealFloat64(a, b, c, d, e float64) []float64 {
	return Float64s(
		RealComponents(
			FilterToReals(
				FirstOrder(
					FromFloat64UseWithCaution(a),
					FromFloat64UseWithCaution(b)))))
}
