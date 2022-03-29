package main

import (
	"fmt"
)

func main() {
	fmt.Print("My ") // this will be executed first
    defer fmt.Print("Go") // this will be defered until the next underfered function is executed. It is put on the bottom of the defer execution stack
    defer fmt.Print("is ") // this will be put on top of the stack
	fmt.Print("name ") // this will be executed second
}