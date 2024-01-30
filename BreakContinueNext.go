package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	arr := make([]int, x+1)
	for i := 0; i < x+1; i++ {
		arr[i] = i
		if arr[i]%10 != 0 {
			if arr[i] < 42 {
				fmt.Println(arr[i])
			} else {
				fmt.Println("ERROR")
				break
			}
		}
	}
}
