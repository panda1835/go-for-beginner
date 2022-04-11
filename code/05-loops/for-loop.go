package main

import (
	"fmt"
)

func main() {
	// calculate sum of odd numbers
	sum := 0
	for i := 1; i < 20; i = i + 2 {
		sum += i
		fmt.Println(sum)
	}
}
