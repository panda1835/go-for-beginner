package main

import "fmt"

func main() {
	var arr [10]string

	arr[3] = "hello"
	fmt.Println("arr is", arr)
	fmt.Println("arr[0] is", arr[0])
	fmt.Println("arr[3] is", arr[3])
	fmt.Println("len(arr) is", len(arr))

	// this statement will cause error because arr is an array of string
	// arr[0] = 1
}
