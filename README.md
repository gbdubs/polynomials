# Precise Polynomial Solver

This package allows you to solve degree four polynomials in constant time with customizable precision.

## How to use it

Shown below are instructions for how to use this for fourth-order polynomials, but the syntax is 
analogous for first, second and third order polynomials.

### Real `big.Float` Roots

```
// Only improves on float64 if you have existing precision from elsewhere. 
a := big.NewFloat(123.342) // b := ...; ... e

realRoots := FourthOrderReal(a, b, c, d, e)
// realRoots is an array of zero to four big.Float, each representing a unique root.
// If empty there are no real roots, only complex ones.
// Each member in realRoots will satisfy the equation ax^4 + bx^3 + cx^2 + dx^1 + e = 0
```

### Real Float64 Roots

```
a := 1.24; b := -32.0; c := 55.22; d := -15.0; e := 22.22

realRoots := FourthOrderRealFloat64(a, b, c, d, e)
// realRoots is an array of zero to four float64s, each representing a unique root.
// If empty there are no real roots, only complex ones.
// Each member in realRoots will satisfy the equation ax^4 + bx^3 + cx^2 + dx^1 + e = 0
```

### Complex Roots

```
a := complex(1.3, -3.202) // b := ...

complexRoots := FourthOrderComplex128(a, b, c, d, e)
// complexRoots is an array containing between 1 and 4 complex128s, each representing a unique root. 
// Each member in complexRoots will satisfy the equation ax^4 + bx^3 + cx^2 + dx^1 + e = 0
```

### Arbtitrary Precision

```
a := &bigcomplex.BigComplex{ Real: big.NewFloat(23.2), Imag: big.NewFloat(-12) } 
// ... b, c, d, e

preciceComplexRoots := FourthOrder(a, b, c, d, e)
// complexRoots is an array containing between 1 and 4 *bigcomplex.BigComplex, 
// each representing a unique root. 
// Each member in complexRoots will satisfy the equation ax^4 + bx^3 + cx^2 + dx^1 + e = 0
```

## Precision

You can set the value you want for mantissa precision by calling

```
precision.Set(/* number of bits in the mantissa */ 1000)
```

The default precision is 1000. For context, the precision of a float64 is 53. 

If you care about results that are accurate to the floating point level, I don't recommend using a
precision less than 400, since that can start to bleed actual error into your end results (when viewed
as a float64).

Shown below are the actual errors when evaluating the roots through the polynomial that was given,
and evaluating the magnitude of the error at the 1000 level. Note that it is incredibly unlikely to
get an error with an exponent in the 40-50 range, but not impossible. That's because of the compounding
trailing errors from over 10,000 simple operations.

![A descritpion of the exponentially weighted error, which demonstrates that almost all of the results have no error, but a long tail have error on the order of -950 on the exponent
for a precision of 1000](./img/1000-precision.png)

## Runtime

Runtime increases with precision. For the fourth-order calculations,

* 400 precision = 4.68ms 
* 500 precision = 5.05ms 
* 1000 precision = 8.34ms 
* 2000 precision = 21.3ms 
* 5000 precision = 150ms 

Unless you're using more precice inputs, 400 precision is totally fine for float64 level precision.

## How it works 

The library leans heavily on others' work to do floating point operations at arbitrary precision:

- The golang native `big.Float` libraries
- A (self made) BigComplex struct with fast implementations for common operations including `Sqrt`
- Shims of [ivy](https://github.com/robpike/ivy) for some of the more involved/expensive computations
- Implementations from [this paper](https://www.researchgate.net/publication/361483599_Fast_Trigonometric_functions_for_Arbitrary_Precision_number) for faster trigonometric identities than ivy provides.

It also relies on some standard equation sets for computing roots of polynomials, including:

- Wikipedia's [Generalized Cubic Formula](https://en.wikipedia.org/wiki/Cubic_equation#General_cubic_formula)
- A [solution to cubics](https://en.wikipedia.org/wiki/Cubic_equation#General_cubic_formula) from the University of Vanderbilt
- and [one of the most helpful stack exchange posts I've ever seen, on the solution to a fourth order polynomial](https://math.stackexchange.com/a/786/1072683)

## Bugs

Please feel free to file a bug if you find one, I'm motivated to keep this library well supported.