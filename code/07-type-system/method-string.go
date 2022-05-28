package main

import "fmt"

type Car struct {
	color string
	brand string
}

func (c *Car) changeColor(new_color string) {
	c.color = new_color
}

func (c Car) String() string {
	return fmt.Sprintf("Brand: %s, Color: %s", c.brand, c.color)
}

func main() {
	BMW := Car{
		color: "blue",
		brand: "BMW",
	}

	fmt.Println("Before change")
	fmt.Println(BMW)

	BMW.changeColor("red")

	fmt.Println("After change")
	fmt.Println(BMW)

}
