package main

import "fmt"

func main() {

	result := 1
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			for k := 0; k < 300; k++ {
				for l := 0; l < 300; l++ {
					result = result + 1
				}
			}
		}
	}

	fmt.Println(result)
}
