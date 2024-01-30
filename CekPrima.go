package main

import "fmt"

func cekprima(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if cekprima(i) {
			fmt.Println("YA")
		} else {
			fmt.Println("BUKAN")
		}
	}
}
