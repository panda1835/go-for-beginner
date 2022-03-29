package main

import (
	"fmt"
)

func main() {
	y := 3
	if y >= 5 {
		fmt.Println(5)
	} else if y < 2 {
		fmt.Println(2)
	} else {
		fmt.Println(y)
	}
}