package main

import "fmt"

func main() {
	odds := [10]int{1, 3, 5, 7, 9}

	var odd []int = odds[0:3]

	fmt.Println(odd)
}
