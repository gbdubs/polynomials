package polynomials

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/gbdubs/polynomials/bigcomplex"
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

func randomBigComplex() *bigcomplex.BigComplex {
	return &bigcomplex.BigComplex{
		Real: newFloat().SetFloat64(rand.NormFloat64()),
		Imag: newFloat().SetFloat64(rand.NormFloat64()),
	}
}

func TestFirstOrder(t *testing.T) {
	nTestsToRun := 100
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		a := randomBigComplex()
		b := randomBigComplex()
		roots := FirstOrder(a, b)
		if len(roots) < 1 || len(roots) > 1 {
			t.Fatalf("expected [1, 1] solutions, got [%d]: %+v", len(roots), roots)
		}
		for _, x := range roots {
			assertBigComplexEq(t, add(mul(a, x), b), fromInt(0))
		}
		assertUnique(t, roots)
	}
}

func TestSecondOrder(t *testing.T) {
	nTestsToRun := 100
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		a := randomBigComplex()
		b := randomBigComplex()
		c := randomBigComplex()
		roots := SecondOrder(a, b, c)
		if len(roots) < 1 || len(roots) > 2 {
			t.Fatalf("expected [1, 2] solutions, got [%d]: %+v", len(roots), roots)
		}
		for _, x := range roots {
			assertBigComplexEq(t, add(mul(a, x, x), mul(b, x), c), fromInt(0))
		}
		assertUnique(t, roots)
	}
}

func TestThirdOrder(t *testing.T) {
	nTestsToRun := 100
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		a := randomBigComplex()
		b := randomBigComplex()
		c := randomBigComplex()
		d := randomBigComplex()
		fmtFloatForWA := func(f float64) string {
			sign := ""
			if f > 0 {
				sign = "+"
			}
			return fmt.Sprintf("%s%.16f", sign, f)
		}
		fmtComplexForWA := func(bc *bigcomplex.BigComplex) string {
			r, _ := bc.Real.Float64()
			i, _ := bc.Imag.Float64()
			return fmt.Sprintf("(0 %s %si)", fmtFloatForWA(r), fmtFloatForWA(i))
		}
		desc := fmt.Sprintf("%sx^3 + %sx^2 + %sx + %s = 0\n",
			fmtComplexForWA(a),
			fmtComplexForWA(b),
			fmtComplexForWA(c),
			fmtComplexForWA(d))
		t.Run(desc, func(t *testing.T) {
			roots := ThirdOrder(a, b, c, d)
			if len(roots) < 1 || len(roots) > 3 {
				t.Fatalf("expected [1, 3] solutions, got [%d]: %+v", len(roots), roots)
			}
			for _, x := range roots {
				assertBigComplexEq(t, add(mul(a, x, x, x), mul(b, x, x), mul(c, x), d), fromInt(0))
			}
			assertUnique(t, roots)
		})
	}
}

func TestFourthOrder(t *testing.T) {
	nTestsToRun := 100
	rand.Seed(0)
	for i := 0; i < nTestsToRun; i++ {
		a := randomBigComplex()
		b := randomBigComplex()
		c := randomBigComplex()
		d := randomBigComplex()
		e := randomBigComplex()
		roots := FourthOrder(a, b, c, d, e)
		if len(roots) < 1 || len(roots) > 4 {
			t.Fatalf("expected [1, 4] solutions, got [%d]: %+v", len(roots), roots)
		}
		for _, x := range roots {
			assertBigComplexEq(t, add(mul(a, x, x, x, x), mul(b, x, x, x), mul(c, x, x), mul(d, x), e), fromInt(0))
		}
		assertUnique(t, roots)
	}
}

func assertBigComplexEq(t *testing.T, got, want *bigcomplex.BigComplex) {
	t.Helper()
	assertBigEq(t, "real", got.Real, want.Real)
	assertBigEq(t, "imag", got.Imag, want.Imag)
}

func assertBigEq(t *testing.T, context string, got, want *big.Float) {
	t.Helper()
	comparer := cmp.Comparer(func(a, b *big.Float) bool {
		aa := newFloat().Set(a)
		bb := newFloat().Set(b)
		c := aa.Sub(aa, bb)
		c = c.Abs(c)
		fmt.Printf("Diff Magnitude = %+v\n", c.MantExp(nil))
		return c.Cmp(getCmpTolerance()) < 0
	})
	if diff := cmp.Diff(want, got, comparer); diff != "" {
		t.Fatalf("%s: got %+v, want %+v, diff %s", context, got, want, diff)
	}
}

func assertUnique(t *testing.T, bcs []*bigcomplex.BigComplex) {
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
