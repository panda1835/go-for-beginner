# 8. Inheritance

## Interface
Interfaces are types that declares behaviours [1]. The behaviours are not implemented by the interface but by the type that inherits it. The implementation is done through methods. In Go, implementation of interface is implicit, which means that there is no clear signal of a type implements an interface. Rather, it is done explicitly by the type that overrides all methods declared by the interface. And by Liskov substitution principle, an object of a type can be substituted with the interface that the type implicitly implements.

Syntax
```text
type <interface name> interface{
    <method 1>() <return type> (optional)
    ...
}
```
Example
```go
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
```
In this example, Square and Rectangle implements 2 methods of Shape so it implicitly implements the Shape interface. `print()` function uses a Shape object as its parameter. Thus, objects of Square and Rectangle can be fed into `print()` function. However, Circle only implements one of Shape methods, which means that Circle is not an implementation of Shape interface. Thus, it can not be used in `print()` function.

## Reference
- Go in Action - Chapter 5.4