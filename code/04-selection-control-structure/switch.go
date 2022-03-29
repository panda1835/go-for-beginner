package main

import (
	"fmt"
)

func main() {
	y := 3
	switch y {
        case 1:
            fmt.Println("one")
            fmt.Println(1)
        case 2:
            fmt.Println("two")
            fmt.Println(2)
        default:
            fmt.Println("three")
            fmt.Println(3)
        
    }
}