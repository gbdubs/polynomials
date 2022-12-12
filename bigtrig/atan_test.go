package bigtrig

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAtan(t *testing.T) {
	cases := []struct {
		val float64
	}{
		{val: 1.234},
		{val: -1.234},
		{val: 0},
		{val: 10000000},
		{val: -0.222222},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%f", c), func(t *testing.T) {
			expected := math.Atan(c.val)
			actual, _ := Atan(newFloat().SetFloat64(c.val)).Float64()
			if diff := cmp.Diff(expected, actual, float64Comparison()); diff != "" {
				t.Errorf("unexpected diff:\n%s", diff)
			}
		})
	}
}

func float64Comparison() cmp.Option {
	return cmpopts.EquateApprox(
		.0000000000001,
		.0000000000001)
}
