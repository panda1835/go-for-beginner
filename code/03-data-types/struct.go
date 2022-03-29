package main

import (
	"fmt"
)

type student struct {
	name string
	GPA  float64
}

func main() {
	studentA := student{name: "A", GPA: 3.8}
	fmt.Println(studentA)
	fmt.Println(studentA.name)
	fmt.Println(studentA.GPA)
}
