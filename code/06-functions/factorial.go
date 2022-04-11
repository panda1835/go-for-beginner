package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return n * factorial(n-1)
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
}
