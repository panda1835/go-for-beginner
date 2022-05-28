package main

import "fmt"

type Car struct {
	color string
	brand string
}

func (c *Car) changeColorWithPointerReceiver(new_color string) {
	c.color = new_color
}

func (c Car) changeColorWithValueReceiver(new_color string) {
	c.color = new_color
}

func main() {
	BMW := Car{
		color: "blue",
		brand: "BMW",
	}

	fmt.Println("Before change")
	fmt.Println(BMW)
	fmt.Println("======")
	fmt.Println("After applying change not using pointer")
	BMW.changeColorWithValueReceiver("red")
	fmt.Println(BMW)

	fmt.Println("After applying change using pointer")
	BMW.changeColorWithPointerReceiver("red")
	fmt.Println(BMW)
}
