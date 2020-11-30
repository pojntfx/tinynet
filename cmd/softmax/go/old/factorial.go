package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(10))
}

func factorial(n int) int {

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
