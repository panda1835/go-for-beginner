package main

import (
	"fmt"
)

type Student struct {
	name string
	id   string
}

type StudentRecord struct {
	student  Student
	semester string
	grade    float32
}

func main() {
	spring_sem := StudentRecord{Student{"A", "001"}, "spring", 3.8}
	fall_sem := StudentRecord{
		student: Student{
			name: "A",
			id:   "001",
		},
		grade:    3.78,
		semester: "fall",
	}
	fmt.Println("student", spring_sem.student.name, "semester", spring_sem.semester, "grade", spring_sem.grade)
	fmt.Println("student", fall_sem.student.name, "semester", fall_sem.semester, "grade", fall_sem.grade)
}
