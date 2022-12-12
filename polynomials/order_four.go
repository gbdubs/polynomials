package polynomials

import (
	"math/big"

	"github.com/gbdubs/polynomials/bigcomplex"
)

// Thank goodness for math stack exchange
// https://math.stackexchange.com/a/786/1072683
func FourthOrder(a, b, c, d, e *bigcomplex.BigComplex) []*bigcomplex.BigComplex {
	if a.IsZero() {
		return ThirdOrder(b, c, d, e)
	}

	p1 := add(
		mul(fromInt(2), c, c, c),
		mul(fromInt(-9), b, c, d),
		mul(fromInt(27), a, d, d),
		mul(fromInt(27), b, b, e),
		mul(fromInt(-72), a, c, e))
	p2 := add(
		p1,
		sqrt(add(
			mul(
				fromInt(-4),
				cube(
					add(
						sq(c),
						mul(fromInt(-3), b, d),
						mul(fromInt(12), a, e)))),
			sq(p1),
		)),
	)
	p3 := add(
		div(
			add(
				sq(c),
				mul(fromInt(-3), b, d),
				mul(fromInt(12), a, e)),
			mul(fromInt(3), a, cbrt(div(p2, fromInt(2))))),
		div(
			cbrt(div(p2, fromInt(2))),
			mul(fromInt(3), a)))

	if p2.IsZero() {
		p3 = fromInt(0)
	}
	p4 := sqrt(add(
		div(sq(b), mul(fromInt(4), a, a)),
		div(mul(fromInt(-2), c), mul(fromInt(3), a)),
		p3))
	p5 := add(
		div(sq(b), mul(fromInt(2), a, a)),
		div(mul(fromInt(-4), c), mul(fromInt(3), a)),
		mul(fromInt(-1), p3))
	p6 := div(
		add(
			div(mul(fromInt(-1), cube(b)), cube(a)),
			div(mul(fromInt(4), b, c), sq(a)),
			div(mul(fromInt(-8), d), a)),
		mul(fromInt(4), p4))
	if p4.IsZero() {
		p6 = fromInt(0)
	}
	p7 := div(mul(fromInt(-1), b), mul(fromInt(4), a))
	p8 := div(p4, fromInt(2))
	p9 := div(sqrt(sub(p5, p6)), fromInt(2))
	p10 := div(sqrt(add(p5, p6)), fromInt(2))
	xs := []*bigcomplex.BigComplex{}
	if p9.IsZero() {
		xs = append(xs, sub(p7, p8))
	} else {
		xs = append(xs, sub(sub(p7, p8), p9))
		xs = append(xs, add(sub(p7, p8), p9))
	}
	if p10.IsZero() {
		xs = append(xs, add(p7, p8))
	} else {
		xs = append(xs, sub(add(p7, p8), p10))
		xs = append(xs, add(p7, p8, p10))
	}
	return xs
}

func FourthOrderComplex128(a, b, c, d, e complex128) []complex128 {
	return ToComplex128s(
		FourthOrder(
			FromComplex128UseWithCaution(a),
			FromComplex128UseWithCaution(b),
			FromComplex128UseWithCaution(c),
			FromComplex128UseWithCaution(d),
			FromComplex128UseWithCaution(e)))
}

func FourthOrderReal(a, b, c, d, e *big.Float) []*big.Float {
	return RealComponents(
		FilterToReals(
			FourthOrder(
				&bigcomplex.BigComplex{Real: newFloat().Set(a), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(b), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(c), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(d), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(e), Imag: newFloat()})))
}

func FourthOrderRealFloat64(a, b, c, d, e float64) []float64 {
	return Float64s(
		RealComponents(
			FilterToReals(
				FourthOrder(
					FromFloat64UseWithCaution(a),
					FromFloat64UseWithCaution(b),
					FromFloat64UseWithCaution(c),
					FromFloat64UseWithCaution(d),
					FromFloat64UseWithCaution(e)))))
}
