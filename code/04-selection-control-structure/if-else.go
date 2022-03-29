package main

import (
	"fmt"
)

func main() {
    x := 10
	if y := x - 5; y >= 0 { 
		fmt.Println(y)
	} else{
        fmt.Println(-y)
    }
}