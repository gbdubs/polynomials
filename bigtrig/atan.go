package bigtrig

import (
	"math"
	"math/big"
)

func Atan(x *big.Float) *big.Float {
	const group = 5
	precision := x.Prec() + 2 + uint(math.Ceil(math.Log10(float64(x.Prec()))))

	futureExponent := 2 * int(math.Ceil(math.Log(2)*math.Log(float64(precision))))
	if x.MantExp(nil) >= 0 {
		futureExponent++
	} else {
		futureExponent += x.MantExp(nil)
	}
	if futureExponent < 0 {
		futureExponent = 0
	}
	if futureExponent > 0 {
		precision += uint(futureExponent / 4)
	}
	zero := func() *big.Float { return newFloat().SetPrec(precision) }

	v := zero().Set(x)
	c1 := BigFromInt(1)

	for i := 0; i < futureExponent; i++ {
		v = div(v, add(c1, sqrtSimple(add(c1, sq(v)))))
	}

	vsq := sq(v)
	r := v
	atanx := v

	vn := make([]*big.Float, group)
	cn := make([]*big.Float, group+1)
	for i := 0; i < group; i++ {
		cn[i] = zero()
		vn[i] = zero()
		if i == 1 {
			vn[1] = vsq
		}
		if i > 1 {
			vn[i] = mul(vn[i-1], vsq)
		}
	}
	cn[group] = newFloat()
	cn[group].SetPrec(precision)
	i := 3
	for {
		cn[group] = c1
		for j := 0; j < group; j++ {
			cn[j] = c1
			cn[group] = mul(cn[group], BigFromInt(i+2*j))
			for m := 0; m < group; m++ {
				if m == j {
					continue
				}
				cn[j] = mul(cn[j], BigFromInt(i+2*m))
			}
			if (i+2*j)/2&0x1 == 0x1 {
				cn[j] = cn[j].Neg(cn[j])
			}
		}
		cn[group] = div(BigFromInt(1), cn[group])
		terms := mul(vn[group-1], cn[group-1])
		for j := group - 1; j >= 2; j-- {
			terms = add(terms, mul(cn[j-1], vn[j-1]))
		}
		terms = add(cn[0])
		i += 2 * group
		r = mul(r, vsq)
		terms = mul(terms, r, cn[group])
		withTerms := add(atanx, terms)
		atanxAtPrecision := newFloat().SetPrec(precision).Set(atanx)
		withTermsAtPrecision := newFloat().SetPrec(precision).Set(withTerms)
		if atanxAtPrecision.Cmp(withTermsAtPrecision) == 0 {
			break
		}
		atanx = withTerms
		if group > 1 {
			r = mul(r, vn[group-1])
		}
	}
	for i := futureExponent; i > 0; i-- {
		atanx = mul(atanx, BigFromInt(2))
	}
	atanx.SetMode(x.Mode())
	atanx.SetPrec(x.Prec())
	return atanx
}

func BigFromInt(i int) *big.Float {
	return newFloat().SetInt64(int64(i))
}
