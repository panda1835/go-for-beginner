package main

import (
	"fmt"
)

func split_string(a string, i int) (a1, a2 string) {
	// this function splits a string at a specific index
	a1, a2 = a[:i], a[i:]
	return
}

func main() {
	var s = "Hello World!"
	var i = 4
	var s1, s2 = split_string(s, i)
	fmt.Printf("String '%s' is divided at %d into '%s' and '%s'", s, i, s1, s2)
}
