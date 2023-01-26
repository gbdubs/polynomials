package maybebigcomplex

import (
	"math/big"

	"github.com/gbdubs/maybebig/maybebig"
	"github.com/gbdubs/polynomials/ivyshims"
)

type BigComplex struct {
	Real *maybebig.Float
	Imag *maybebig.Float
}

func (bc *BigComplex) Magnitude() *maybebig.Float {
	return sqrtSimple(add(sq(bc.Real), sq(bc.Imag)))
}

func (bc *BigComplex) Theta() *maybebig.Float {
	if isZero(bc.Real) {
		if ltz(bc.Imag) {
			return div(maybebig.Pi(), fromInt(-2))
		}
		return div(maybebig.Pi(), fromInt(2))
	}
	atan := fromBig(ivyshims.Atan(div(bc.Imag, bc.Real).GetBig()))
	if ltz(bc.Real) {
		if gtz(bc.Imag) {
			atan = add(atan, maybebig.Pi())
		} else {
			atan = sub(atan, maybebig.Pi())
		}
	}
	return atan
}

func Add(bcs ...*BigComplex) *BigComplex {
	r := newFloat()
	i := newFloat()
	for _, bc := range bcs {
		r = add(r, bc.Real)
		i = add(i, bc.Imag)
	}
	return &BigComplex{
		Real: r,
		Imag: i,
	}
}

func Sub(a, b *BigComplex) *BigComplex {
	return &BigComplex{
		Real: sub(a.Real, b.Real),
		Imag: sub(a.Imag, b.Imag),
	}
}

func Mul(bcs ...*BigComplex) *BigComplex {
	result := &BigComplex{Real: fromInt(1), Imag: fromInt(0)}
	for _, bc := range bcs {
		result = Product(result, bc)
	}
	return result
}

func Product(a, b *BigComplex) *BigComplex {
	return &BigComplex{
		Real: sub(mul(a.Real, b.Real), mul(a.Imag, b.Imag)),
		Imag: add(mul(a.Real, b.Imag), mul(a.Imag, b.Real)),
	}
}

func Div(bc1, bc2 *BigComplex) *BigComplex {
	a := bc1.Real
	b := bc1.Imag
	c := bc2.Real
	d := bc2.Imag
	return &BigComplex{
		Real: div(add(mul(a, c), mul(b, d)), add(sq(c), sq(d))),
		Imag: div(sub(mul(b, c), mul(a, d)), add(sq(c), sq(d))),
	}
}

func Sq(a *BigComplex) *BigComplex {
	return Mul(a, a)
}

func Cube(a *BigComplex) *BigComplex {
	return Mul(a, a, a)
}

/*
	On it's face this is likely wrong - it occasionally produces 4 roots!!!

// Uses a simplified version of this code
// https://www.geeksforgeeks.org/square-root-of-two-complex-numbers/

	func (bc *BigComplex) Sqrt() []*BigComplex {
		result := []*BigComplex{}
		a := bc.Real
		b := bc.Imag
		two := fromInt(2)
		negOne := fromInt(-1)
		magSq := add(sq(a), sq(b))
		x1 := abs(sqrtSimple(div(add(a, sqrtSimple(magSq)), two)))
		if isZero(x1) {
			y1 := sqrtSimple(mul(fromInt(-1), a))
			result = append(result, &BigComplex{Real: x1, Imag: y1})
		} else {
			y1 := div(b, mul(two, x1))
			result = append(result, &BigComplex{Real: x1, Imag: y1})
			x2 := mul(negOne, x1)
			y2 := div(b, mul(two, x2))
			result = append(result, &BigComplex{Real: x2, Imag: y2})
		}

		x3 := div(sub(a, sqrtSimple(magSq)), two)
		if positive(x3) {
			x3 = abs(sqrtSimple(x3))
			y3 := div(b, mul(two, x3))
			result = append(result, &BigComplex{Real: x3, Imag: y3})
			x4 := mul(negOne, x3)
			y4 := div(b, mul(two, x3))
			result = append(result, &BigComplex{Real: x4, Imag: y4})
		}
		return result
	}
*/
func (bc *BigComplex) Sqrt() *BigComplex {
	// Pursues this approach over the ivyshims version, which requires conversion.
	// r, i := ivyshims.ComplexSqrt(bc.Real.GetBig(), bc.Imag.GetBig())
	mag := sqrt(add(sq(bc.Real), sq(bc.Imag)))
	real := sqrt(div(add(mag, bc.Real), fromInt(2)))
	imag := sqrt(div(sub(mag, bc.Real), fromInt(2)))
	if ltz(bc.Imag) {
		imag = mul(imag, fromInt(-1))
	}
	return &BigComplex{
		Real: real,
		Imag: imag,
	}
}

func (bc *BigComplex) Cbrt() *BigComplex {
	return bc.Pow(fromBig(newFloat().GetBig().SetRat(big.NewRat(1, 3))))
}

func (bc *BigComplex) Pow(e *maybebig.Float) *BigComplex {
	r := bc.Magnitude()
	theta := bc.Theta()

	return Mul(
		&BigComplex{
			Real: fromBig(ivyshims.Pow(r.GetBig(), e.GetBig())),
			Imag: newFloat(),
		},
		&BigComplex{
			Real: fromBig(ivyshims.Cos(mul(e, theta).GetBig())),
			Imag: fromBig(ivyshims.Sin(mul(e, theta).GetBig())),
		},
	)
}

func (bc *BigComplex) Equals(other *BigComplex) bool {
	return eq(bc.Real, other.Real) && eq(bc.Imag, other.Imag)
}

func (bc *BigComplex) IsZero() bool {
	other := newFloat()
	return eq(other, bc.Imag) && eq(other, bc.Real)
}

func Unique(in []*BigComplex) []*BigComplex {
	result := []*BigComplex{}
	for _, i := range in {
		found := false
		for _, j := range result {
			if i.Equals(j) {
				found = true
			}
		}
		if !found {
			result = append(result, i)
		}
	}
	return result
}

func (bc *BigComplex) Complex128() complex128 {
	r := bc.Real.Float64()
	i := bc.Imag.Float64()
	return complex(r, i)
}
