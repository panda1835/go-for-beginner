package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float32
	getPerimeter() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	width  float32
	height float32
}

type Circle struct {
	radius float32
}

// let Square implements Shape interface by implements its
// 2 methods getArea() and getPerimeter()
func (s Square) getArea() float32 {
	return s.side * s.side
}

func (s Square) getPerimeter() float32 {
	return s.side * 4
}

// let Rectangle implements Shape interface by implements its
// 2 methods getArea() and getPerimeter()
func (r Rectangle) getArea() float32 {
	return r.height * r.width
}

func (r Rectangle) getPerimeter() float32 {
	return (r.height + r.width) * 2
}

// Circle does not implement Shape because it only implements
// one of its function
func (c Circle) getArea() float32 {
	return math.Pi * c.radius * c.radius
}

func print(s Shape) {
	fmt.Printf("Area: %f, Diameter: %f\n", s.getArea(), s.getPerimeter())
}

func main() {

	square := Square{
		side: 4,
	}

	rectangle := Rectangle{
		width:  5,
		height: 6,
	}

	// circle := Circle{
	// 	radius: 6,
	// }

	print(square)
	print(rectangle)
	// print(circle) --> error because Circle does not implement Shape interface
}
