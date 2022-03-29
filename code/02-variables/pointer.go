package main

import (
	"fmt"
)

func main() {
	var p *int // initialize pointer variable
	i := 4
	p = &i
	fmt.Println("==Before==")
	fmt.Println("i:", i) // print value of i on its initialization
	fmt.Println("p:", p) // print pointer p (the memory address of i)
	*p = *p / 2          // take the value stored in p and divide it by 2
	fmt.Println("==After==")
	fmt.Println("i:", i) // print p again to see if the memory address changes
	fmt.Println("p:", p) // print value of i to see if its value changes
}
