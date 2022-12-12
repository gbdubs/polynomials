package bigcomplex

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/gbdubs/polynomials/ivyshims"
	"github.com/gbdubs/polynomials/precision"
	"github.com/google/go-cmp/cmp"
)

var cmpTolerance *big.Float

func getCmpTolerance() *big.Float {
	if cmpTolerance != nil {
		return cmpTolerance
	}
	// 10^100 precision.
	b, _, err := big.ParseFloat("0.0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001", 10, precision.Get(), big.ToPositiveInf)
	if err != nil {
		panic(fmt.Errorf("creating cmp tolerance: %w", err))
	}
	cmpTolerance = b
	return b
}

func randomBigComplex() *BigComplex {
	return &BigComplex{
		Real: newFloat().SetFloat64(rand.NormFloat64()),
		Imag: newFloat().SetFloat64(rand.NormFloat64()),
	}
}

func TestSqrt(t *testing.T) {
	nTestsToRun := 1000
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		n := randomBigComplex()
		sqrts := n.Sqrt()
		if len(sqrts) < 1 || len(sqrts) > 2 {
			t.Fatalf("expected [1, 4] solutions, got [%d]: %+v", len(sqrts), sqrts)
		}
		for _, sqrt := range sqrts {
			assertBigComplexEq(t, Sq(sqrt), n)
		}
		assertUnique(t, sqrts)
	}
}

func TestAddSub(t *testing.T) {
	nTestsToRun := 1000
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		a := randomBigComplex()
		b := randomBigComplex()
		c := Sub(a, b)
		assertBigComplexEq(t, a, Add(c, b))
	}
}

func TestDivMul(t *testing.T) {
	nTestsToRun := 1000
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		a := randomBigComplex()
		b := randomBigComplex()
		c := Div(a, b)
		assertBigComplexEq(t, a, Mul(c, b))
	}
}

func TestThetaAndMag(t *testing.T) {
	nTestsToRun := 1000
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		a := randomBigComplex()
		theta := a.Theta()
		mag := a.Magnitude()
		re := mul(mag, ivyshims.Cos(theta))
		im := mul(mag, ivyshims.Sin(theta))
		assertBigComplexEq(t, a, &BigComplex{Real: re, Imag: im})
	}
}

func assertBigComplexEq(t *testing.T, got, want *BigComplex) {
	t.Helper()
	assertBigEq(t, "real", got.Real, want.Real)
	assertBigEq(t, "imag", got.Imag, want.Imag)
}

func assertBigEq(t *testing.T, context string, got, want *big.Float) {
	t.Helper()
	comparer := cmp.Comparer(func(a, b *big.Float) bool {
		return abs(sub(a, b)).Cmp(getCmpTolerance()) < 0
	})
	if diff := cmp.Diff(want, got, comparer); diff != "" {
		t.Fatalf("%s: got %+v, want %+v, diff %s", context, got, want, diff)
	}
}

func assertUnique(t *testing.T, bcs []*BigComplex) {
	result := make(map[string]bool)
	for _, bc := range bcs {
		real := newFloat().Set(bc.Real).String()
		imag := newFloat().Set(bc.Imag).String()
		s := fmt.Sprintf("R=%s I=%s", real, imag)
		if result[s] {
			t.Errorf("found duplicate: %s", s)
		}
		result[s] = true
	}
}
