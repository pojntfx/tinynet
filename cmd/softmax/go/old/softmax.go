package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0.
	result := []float64{}
	input := []float64{1, 1, 3}

	for i := 0; i < len(input); i++ {
		sum = sum + softmaxSum(input[i])
	}

	for i := 0; i < len(input); i++ {
		result = append(result, softmaxResult(sum, input[i]))
	}

	fmt.Println(result)
}

func softmaxSum(input float64) float64 {
	return math.Exp(input)
}

func softmaxResult(sum float64, input float64) float64 {
	return math.Exp(input) / sum
}
