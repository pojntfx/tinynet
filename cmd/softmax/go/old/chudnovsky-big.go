package main

import (
	"fmt"
	"math/big"
)

func main() {
	result := big.NewFloat(0)
	result.SetPrec(1024)
	sum := big.NewFloat(0)
	sum.SetPrec(1024)

	k := 100
	for i := 0; i <= k; i++ {
		iFloat := big.NewFloat(float64(i))
		iFloat.SetPrec(1024)
		sum = sum.Add(sum, sumElement(iFloat))
	}

	a := big.NewFloat(4270934400)
	a.SetPrec(1024)

	b := big.NewFloat(10005)
	b.SetPrec(1024)

	result = Div(a, Mul(Root(b, 2), sum))
	fmt.Println(result)
}

func sumElement(i *big.Float) *big.Float {
	iAsUint64, _ := i.Uint64()
	threeFloat := big.NewFloat(3)
	threeFloat.SetPrec(1024)
	mulAsUint64, _ := Mul(threeFloat, i).Uint64()

	oneFloat := big.NewFloat(-1)
	oneFloat.SetPrec(1024)

	sixFloat := big.NewFloat(6)
	sixFloat.SetPrec(1024)

	bigFloat := big.NewFloat(13591409)
	bigFloat.SetPrec(1024)

	bigOtherFloat := big.NewFloat(545140134)
	bigOtherFloat.SetPrec(1024)

	bigBigFloat := big.NewFloat(640320)
	bigBigFloat.SetPrec(1024)

	return Mul(
		Pow(
			oneFloat,
			iAsUint64),
		Mul(
			Div(
				Factorial(
					Mul(sixFloat, i)),
				Mul(
					Pow(Factorial(i), 3),
					Factorial(Mul(threeFloat, i)),
				),
			),
			Div(
				Add(
					bigFloat,
					Mul(bigOtherFloat, i)),
				Pow(bigBigFloat, mulAsUint64),
			),
		),
	)
}

// Pow returns a to the power of e
func Pow(a *big.Float, e uint64) *big.Float {
	result := Zero().Copy(a)
	if e == 0 {
		one := big.NewFloat(1)
		one.SetPrec(1024)
		return one
	}
	for i := uint64(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

// Root returns the n-th root of a
func Root(a *big.Float, n uint64) *big.Float {
	limit := Pow(New(2), 1024)
	n1 := n - 1
	n1f, rn := New(float64(n1)), Div(New(1.0), New(float64(n)))
	x, x0 := New(1.0), Zero()
	_ = x0
	for {
		potx, t2 := Div(New(1.0), x), a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 = Mul(t2, potx)
			}
			potx = Mul(potx, potx)
		}
		x0, x = x, Mul(rn, Add(Mul(n1f, x), t2))
		if Lesser(Mul(Abs(Sub(x, x0)), limit), x) {
			break
		}
	}
	return x
}

// Abs returns the absolute value of a
func Abs(a *big.Float) *big.Float {
	return Zero().Abs(a)
}

// New creates a new "bigFloat"
func New(f float64) *big.Float {
	r := big.NewFloat(f)
	r.SetPrec(1024)
	return r
}

// Div divides a and b
func Div(a, b *big.Float) *big.Float {
	return Zero().Quo(a, b)
}

// Zero creates a new "bigFloat" with the value 0.0
func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(1024)
	return r
}

// Mul multiplies a and b
func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}

// Add adds a and b
func Add(a, b *big.Float) *big.Float {
	return Zero().Add(a, b)
}

// Sub substracts a and b
func Sub(a, b *big.Float) *big.Float {
	return Zero().Sub(a, b)
}

// Lesser returns if x is lesser than y
func Lesser(x, y *big.Float) bool {
	return x.Cmp(y) == -1
}

// Factorial returns the factorial of n
func Factorial(n *big.Float) *big.Float {

	twoFloat := big.NewFloat(2)
	twoFloat.SetPrec(1024)

	oneFloat := big.NewFloat(1)
	oneFloat.SetPrec(1024)

	if Lesser(n, twoFloat) {
		return oneFloat
	}

	return Mul(n, Factorial(Sub(n, oneFloat)))
}
