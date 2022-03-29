package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["ten"] = 10
	m["one"] = 1
	fmt.Println(m["ten"])
}
