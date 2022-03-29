package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var studentName string
	var studentAge int
	var studentGPA float32
	var isFemale bool
	var complexNum complex128

	studentName = "Student A"
	studentAge = 19
	studentGPA = 3.87
	isFemale = true
	complexNum = cmplx.Sqrt(-5 + 12i)

	fmt.Println(studentName, studentAge, studentGPA, isFemale, complexNum)

	var nonInt int
	var nonBool bool
	var nonString string
	fmt.Println(nonInt, nonBool, nonString)
}
