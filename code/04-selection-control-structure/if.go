package main

import (
	"fmt"
)

func absolute(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	fmt.Println(absolute(5))
	fmt.Println(absolute(0))
	fmt.Println(absolute(-5))
}
