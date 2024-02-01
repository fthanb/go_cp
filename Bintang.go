package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		for j := 0; j < N-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
