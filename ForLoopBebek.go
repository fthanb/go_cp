package main

import (
	"fmt"
)

func main() {
	var N int
	var sum int
	fmt.Scanln(&N)

	total := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &total[i])
		sum += total[i]
	}
	fmt.Println(sum)
}
