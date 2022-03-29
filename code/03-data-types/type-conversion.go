package main

import (
	"fmt"
	 "reflect"
)

func main() {
	var a = 5
	var b float64 = float64(a/2)
	fmt.Println("a =", reflect.TypeOf(a))
	fmt.Println("b =", reflect.TypeOf(b))
}