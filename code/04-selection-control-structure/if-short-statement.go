package main

import (
	"fmt"
)

// this function calculate |x-5|
func absolute_x_minus_5(x int) int {
	if y := x - 5; y >= 0 { // here it evaluates y = x-5 first and compare it with 0
		return y
	}
	return 5 - x
}

func main() {
	fmt.Println(absolute_x_minus_5(5))
	fmt.Println(absolute_x_minus_5(0))
	fmt.Println(absolute_x_minus_5(-5))
}