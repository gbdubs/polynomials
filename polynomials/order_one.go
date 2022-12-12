package polynomials

import "github.com/gbdubs/polynomials/bigcomplex"

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

func FirstOrderRealsFloat64(a, b float64) []float64 {
	return RealComponents(
		FilterToReals(
			FirstOrder(
				FromFloat64UseWithCaution(a),
				FromFloat64UseWithCaution(b))))
}
