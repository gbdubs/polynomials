package polynomials

import (
	"math/big"

	"github.com/gbdubs/maybebig/maybebig"
	"github.com/gbdubs/polynomials/bigcomplex"
	"github.com/gbdubs/polynomials/maybebigcomplex"
)

func FromFloat64UseWithCaution(f float64) *bigcomplex.BigComplex {
	return &bigcomplex.BigComplex{
		Real: newFloat().SetFloat64(f),
		Imag: newFloat(),
	}
}

func FromComplex128UseWithCaution(c complex128) *bigcomplex.BigComplex {
	return &bigcomplex.BigComplex{
		Real: newFloat().SetFloat64(real(c)),
		Imag: newFloat().SetFloat64(real(c)),
	}
}

func ToComplex128s(bcs []*bigcomplex.BigComplex) []complex128 {
	result := make([]complex128, len(bcs))
	for i, bc := range bcs {
		real, _ := bc.Real.Float64()
		imag, _ := bc.Imag.Float64()
		result[i] = complex(real, imag)
	}
	return result
}

func RealComponents(bcs []*bigcomplex.BigComplex) []*big.Float {
	result := make([]*big.Float, len(bcs))
	for i, bc := range bcs {
		result[i] = bc.Real
	}
	return result
}

func RealComponentsMB(bcs []*maybebigcomplex.BigComplex) []*maybebig.Float {
	result := make([]*maybebig.Float, len(bcs))
	for i, bc := range bcs {
		result[i] = bc.Real
	}
	return result
}

func Float64s(bcs []*big.Float) []float64 {
	result := make([]float64, len(bcs))
	for i, bc := range bcs {
		f, _ := bc.Float64()
		result[i] = f
	}
	return result
}

func FilterToReals(bcs []*bigcomplex.BigComplex) []*bigcomplex.BigComplex {
	result := []*bigcomplex.BigComplex{}
	for _, bc := range bcs {
		if bc.Imag.Cmp(big.NewFloat(0)) == 0 {
			result = append(result, bc)
		}
	}
	return result
}

func FilterToRealsMB(bcs []*maybebigcomplex.BigComplex) []*maybebigcomplex.BigComplex {
	result := []*maybebigcomplex.BigComplex{}
	for _, bc := range bcs {
		if maybebig.Eqz(bc.Imag) {
			result = append(result, bc)
		}
	}
	return result
}
