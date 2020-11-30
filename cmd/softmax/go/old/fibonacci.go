package main

import "fmt"

func main() {
	fmt.Println(fibonacci(100))
}

func fibonacci(n int) int {

	a := 1
	b := 1
	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + b

	}

	return b
}
