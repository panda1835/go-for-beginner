package main

import (
	"fmt"
)

func isFalse() bool {
    fmt.Println("isFalse is executed")
    return false
}

func isTrue() bool {
    fmt.Println("isTrue is executed")
    return true
}

func main() {
	if isFalse() && isTrue(){
        fmt.Println("No short circuit")
    } 
}