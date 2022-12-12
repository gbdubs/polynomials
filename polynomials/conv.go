package polynomials

import (
	"math/big"

	"github.com/gbdubs/polynomials/bigcomplex"
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

func RealComponents(bcs []*bigcomplex.BigComplex) []float64 {
	result := make([]float64, len(bcs))
	for i, bc := range bcs {
		real, _ := bc.Real.Float64()
		result[i] = real
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
