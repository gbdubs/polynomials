package polynomials

import (
	"math/big"

	"github.com/gbdubs/polynomials/bigcomplex"
)

func ThirdOrder(a, b, c, d *bigcomplex.BigComplex) []*bigcomplex.BigComplex {
	if a.IsZero() {
		return SecondOrder(b, c, d)
	}

	// Thank you Wikipedia!
	// https://en.wikipedia.org/wiki/Cubic_equation#General_cubic_formula
	delta0 := sub(sq(b), mul(fromInt(3), a, c))
	delta1 := add(
		mul(fromInt(2), cube(b)),
		mul(fromInt(-9), a, b, c),
		mul(fromInt(27), a, a, d))
	bigC := fromInt(0)
	delta0OverBigC := fromInt(0)
	if !delta0.IsZero() || !delta1.IsZero() {
		innerRoot := sqrt(sub(sq(delta1), mul(fromInt(4), cube(delta0))))
		bigC = cbrt(div(
			add(delta1, innerRoot),
			fromInt(2),
		))
		if bigC.IsZero() {
			bigC = cbrt(div(
				sub(delta1, innerRoot),
				fromInt(2),
			))
		}
		delta0OverBigC = div(delta0, bigC)
	}
	zeta := div(add(fromInt(-1), fromInt(-3).Sqrt()), fromInt(2))
	zeta1 := zeta
	zeta2 := mul(zeta, zeta)
	x0 := div(add(b, bigC, delta0OverBigC), mul(fromInt(-3), a))
	x1 := div(add(b, mul(zeta1, bigC), div(delta0OverBigC, zeta1)), mul(fromInt(-3), a))
	x2 := div(add(b, mul(zeta2, bigC), div(delta0OverBigC, zeta2)), mul(fromInt(-3), a))
	return bigcomplex.Unique([]*bigcomplex.BigComplex{x0, x1, x2})
}

func ThirdOrderEquivalentButSlower(a, b, c, d *bigcomplex.BigComplex) []*bigcomplex.BigComplex {
	if a.IsZero() {
		return SecondOrder(b, c, d)
	}

	// Thank you Vandy!
	// https://math.vanderbilt.edu/schectex/courses/cubic/
	// Gets the first cube root from this formula
	negOne := fromInt(-1)
	three := fromInt(3)
	six := fromInt(6)

	p := div(mul(negOne, b), mul(three, a))
	q := add(mul(mul(p, p), p), div(sub(mul(b, c), mul(three, mul(a, d))), mul(six, sq(a))))
	r := div(c, mul(three, a))
	s := add(sq(q), cube(sub(r, sq(p)))).Sqrt()
	x0 := add(add(q, s).Pow(newRat(1, 3)), sub(q, s).Pow(newRat(1, 3)), p)

	// Then uses polynomial long division (with an assumed remainder of zero)
	// for the other two roots. We then test to see if those roots ACTUALLY
	// satisfy (i.e. if the remainder was actually zero)
	qa := a
	qb := add(b, mul(a, x0))
	qc := add(c, mul(b, x0), mul(a, x0, x0))

	otherRoots := SecondOrder(qa, qb, qc)
	return bigcomplex.Unique(append([]*bigcomplex.BigComplex{x0}, otherRoots...))
}

func ThirdOrderComplex128(a, b, c, d complex128) []complex128 {
	return ToComplex128s(
		ThirdOrder(
			FromComplex128UseWithCaution(a),
			FromComplex128UseWithCaution(b),
			FromComplex128UseWithCaution(c),
			FromComplex128UseWithCaution(d)))
}

func ThirdOrderReal(a, b, c, d *big.Float) []*big.Float {
	return RealComponents(
		FilterToReals(
			ThirdOrder(
				&bigcomplex.BigComplex{Real: newFloat().Set(a), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(b), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(c), Imag: newFloat()},
				&bigcomplex.BigComplex{Real: newFloat().Set(d), Imag: newFloat()})))
}

func ThirdOrderRealFloat64(a, b, c, d, e float64) []float64 {
	return Float64s(
		RealComponents(
			FilterToReals(
				ThirdOrder(
					FromFloat64UseWithCaution(a),
					FromFloat64UseWithCaution(b),
					FromFloat64UseWithCaution(c),
					FromFloat64UseWithCaution(d)))))
}
