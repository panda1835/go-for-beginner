package main

import (
	"fmt"
)

type Student struct {
	name string
	id   string
}

func main() {
	studentA := Student{"A", "001"}
	studentB := Student{
		name: "B",
		id:   "002",
	}
	fmt.Println(studentA)
	fmt.Println(studentB)
}
