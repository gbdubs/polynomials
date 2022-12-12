package polynomials

import (
	"fmt"
	"math"
	"math/cmplx"
)

func ThirdOrderLegacy(a, b, c, d complex128) []complex128 {
	// Thank you Vandy!
	// https://math.vanderbilt.edu/schectex/courses/cubic/
	// Gets the first cube root from this formula
	p := -1 * b / (3 * a)
	q := p*p*p + (b*c-3*a*d)/(6*a*a)
	r := c / (3 * a)
	s := cmplx.Sqrt(q*q + cmplx.Pow(r-(p*p), 3))
	x0 := cmplx.Pow(q+s, 1.0/3.0) + cmplx.Pow(q-s, 1.0/3.0) + p

	// Then uses polynomial long division (with an assumed remainder of zero)
	// for the other two roots. We then test to see if those roots ACTUALLY
	// satisfy (i.e. if the remainder was actually zero)
	qa := a
	qb := b + a*x0
	qc := c + b*x0 + a*x0*x0
	otherRoots := SecondOrderComplex128(qa, qb, qc)
	return append(otherRoots, x0)
}

// Thank goodness for math stack exchange
// https://math.stackexchange.com/a/786/1072683
func FourthOrderLegacy(a, b, c, d, e complex128) []complex128 {
	p1 := 2*c*c*c - 9*b*c*d + 27*a*d*d + 27*b*b*e - 72*a*c*e
	p2 := p1 + customSqrt(-4*cmplx.Pow(c*c-3*b*d+12*a*e, 3)+p1*p1)
	p3 := ((c*c - 3*b*d + 12*a*e) / (3 * a * customCubeRoot(p2/2))) + (customCubeRoot(p2/2) / (3 * a))
	if p2 == complex(0, 0) {
		p3 = complex(0, 0)
	}
	p4 := customSqrt(((b * b) / (4 * a * a)) - ((2 * c) / (3 * a)) + p3)
	p5 := (b*b)/(2*a*a) - (4*c)/(3*a) - p3
	p6 := (((-1 * b * b * b) / (a * a * a)) + ((4 * b * c) / (a * a)) - (8 * d / a)) / (4 * p4)
	if p4 == complex(0, 0) {
		p6 = complex(0, 0)
	}
	p7 := -1 * b / (4 * a)
	p8 := p4 / 2
	p9 := customSqrt(p5-p6) / 2
	p10 := customSqrt(p5+p6) / 2
	xs := []complex128{}
	if isComplexZero(p9) {
		xs = append(xs, p7-p8)
	} else {
		xs = append(xs, p7-p8-p9)
		xs = append(xs, p7-p8+p9)
	}
	if isComplexZero(p10) {
		xs = append(xs, p7+p8)
	} else {
		xs = append(xs, p7+p8-p10)
		xs = append(xs, p7+p8+p10)
	}
	return xs
}

func customSqrt(z complex128) complex128 {
	a := cmplx.Pow(z, .5)
	angle := math.Atan2(imag(a), real(a))
	if angle < math.Pi/-2 {
		panic(fmt.Errorf("unexpected sqrt < -pi/2 = %+v, %+v, %v", z, a, angle))
	}
	if angle > math.Pi/2 {
		panic(fmt.Errorf("unexpected sqrt > pi/2 = %+v, %+v, %v", z, a, angle))
	}
	return a
}

func customCubeRoot(z complex128) complex128 {
	a := cmplx.Pow(z, 1.0/3.0)
	angle := math.Atan2(imag(a), real(a))
	if angle < math.Pi/-3 {
		panic(fmt.Errorf("unexpected cbrt < -pi/3 = %+v, %+v, %v", z, a, angle))
	}
	if angle > math.Pi/3 {
		panic(fmt.Errorf("unexpected cbrt > pi/3 = %+v, %+v, %v", z, a, angle))
	}
	return a
}

func isComplexZero(c complex128) bool {
	return 0 == real(c) && 0 == imag(c) // TODO SOMETHING BETTER HERE, A BUG IS HERE.
}
