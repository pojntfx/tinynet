package main

import (
	"fmt"
	"math"
)

func main() {

	sum := 0.
	k := 28.

	for i := 0; i <= int(k); i++ {
		sum = sum + sumElement(float64(i))
	}

	result := 4270934400. / (math.Sqrt(10005.) * sum)

	fmt.Println(result)
}

func sumElement(k float64) float64 {
	return (math.Pow(-1, k) * (factorial(6*k) / (math.Pow(factorial(k), 3) * factorial(3*k))) * ((13591409 + 545140134*k) / math.Pow(640320, 3*k)))
}

func factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
