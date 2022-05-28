# Type System

## User-defined types
Go does not have classes, but it have types that support users to define their own data type. The most common way to create a new data type is by using `struct`.

Syntax
```text
type <type name> struct {
    ...
    <field name> <field type>
    ...
}
```

Here, `<field type>` can be of any primitive types, or a user-defined type.

There are 2 ways to initialize an object of a new type (if there is no initialization, the object's will be assigned with zero value, i.e., 0 for numeric types, empty for string and false for boolean) [\[1\]].

Syntax
```text
<type name>{<field 1 value>, <field 2 value>, ...}
```
or 
```text
<type name>{
    <field name 1>: <value>,
    <field name 2>: <value>,
    ...
}
```

To retrieve the value of the attribute, simply write the attribute after the object name, seperated by a dot `.`.

Syntax
```text
<object name> = <type initialization>
<object name>.<attribute name>
```

Example
```go
// struct-example.go

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
```
You should receive the output
```text
{A 001}
{B 002}
```

Below is another example of embeded data type in data type. In this example, data type StudentRecord contains a field student of Student type.
```go
// embedding-type.go

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
```
The output should be 
```text
student A semester spring grade 3.8
student A semester fall grade 3.78
```

## Methods
Similar to class of other OOP languages such as Java, type of Go has methods associated to it. To distinguish between a method and a function, a method is a function with a *receiver* argument. A receiver is an argument and it is located between the function definition keyword `func` and the method name. 

Syntax
```text
func (<variable> <user-defined type>) <function name> (parameters) <return type>{
    ...
    <code block>
    ...
    <return statement>
}
```

The receiver is either *pointer* receiver or *value* receiver. For *pointer* receiver, changes applied to object's attributes will be saved. But with *value* receiver, the change is ignored after the method is called. \[1\]

Example
```go
// method-example.go

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

```
You should receive an output
```text
Before change
{blue BMW}
======
After change not using pointer
{blue BMW}
After change using pointer
{red BMW}
```

## `String()` Method
There are standard methods in Go that serves similar purpose across all objects. One of them is `String()` method, which specifies the string representation of an object. \[[2]\]

Syntax
```text
func (<variable> <user-defined type>) String() string{
	return <string representation of the object>
}
```

Example
```go
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
```

This method define a string representation of a Car's object to be in the format: 
```text
Brand: <brand name>, Color: <color>
```
That format is featured in the return value of `String()` method.
The expected output of the program should be
```text
Before change
Brand: BMW, Color: blue
After change
Brand: BMW, Color: red
```

## Reference:
- Go in Action
- [ToString() function in Go][2]

[2]: https://stackoverflow.com/questions/13247644/tostring-function-in-go