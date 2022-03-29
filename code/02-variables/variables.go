package main

import "fmt"

func main() {
	// normal declaration
	var x, y int
	fmt.Println(x, y)
	// with initializers and with type
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	// with initializers and without type
	var m, n, p = 4, "five", false
	fmt.Println(m, n, p)
	// short assigment inside function
	k := true
	fmt.Println(k)

}
