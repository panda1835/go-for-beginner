package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var m, n = 10, 20
	fmt.Println("Before swap m, n =", m, n)

	swap(m, n)

	fmt.Println("After swap m, n =", m, n)
}
