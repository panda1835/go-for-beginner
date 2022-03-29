# 3. Data Types
## 3.1 Is Go statically or dynamically typed?
Go is a strongly and statically typed language, which means that a variable of one type is known at compile time and can not be changed or automatically converted to another type throughout the program execution [\[1\]][1].

Although Go is a strongly typed language, it supports both explicitly and implicitly typed. Go infers the type of untyped variables from its value. The default type of an untyped constant is `bool`, `int`, `float64`, `complex128`, or `string` [\[2\]][2].

```go
// implicit-explicit-typed.go
package main
import (
  "fmt"
  "reflect"
)
  
func main(){
  var a int = 5 
  b := 3.5
  fmt.Println("type of a is", reflect.TypeOf(a))
  fmt.Println("type of b is", reflect.TypeOf(b))
}
```
You will receive the following output for the above code
```text
type of a is int
type of b is float64
```

## 3.2 Basic Data Types
``` go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

When a variable is declared without an explicit initial value, it is given the value that corresponds to `0`.

- Numerics: `0`
- Boolean: `false`
- String: '' (the empty string)

Note: The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. 

``` go
// basic-data-types.go
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
```
Results:
``` text
>> Student A 19 3.87 true (2+3i)
>> 0 false ''
```

Note that you **can not** do operations on variables of different data types (even between `int` and `float`). If you run the following code, it will raise an error says `invalid operation: studentGPA + studentAge (mismatched types float32 and int)`.

```go
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	
	var studentAge int
	var studentGPA float32
	fmt.Println(studentGPA + studentAge)
}
```
For division, if you take `5` and you divide it into `2`, you will get a `2` instead of `2.5`.
However, like any other languages, Go allows you to convert from one basic data type to another, just simply wrap the type name outside the variable. In the following example, `b` is a float conversion of int `a`. We use `reflect` module to print out the type of `a` and `b`.  
```go
// type-conversion.go
package main

import (
	"fmt"
	 "reflect"
)

func main() {
	var a = 5
	var b float64 = float64(a/2)
	fmt.Println("a =", reflect.TypeOf(a))
	fmt.Println("b =", reflect.TypeOf(b))
}
```
You should expect the following output when you run the above code.
```text
a = int
b = float64
```
### 3.3 Operators
#### 3.3.1 Arithmetic Operators
Arithmetic operators apply to numeric values and yield a result of the same type as the first operand. The following list is a summary of all arithmetic operators in Go [\[1\]][1].
```text
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << integer >= 0
>>   right shift            integer >> integer >= 0
```
#### 3.3.2 String Concatenation
Strings can be concatenated using the + operator or the += assignment operator. String addition creates a new string by concatenating the operands [\[1\]][1].
```go
// string-concatenation.go
package main
import (
  "fmt"
)
  
func main(){
    var s1 = "hello"
    var s2 = " world"
    var s3 = s1 + s2
    fmt.Println(s3)
}
```
The above code will print out
```text
hello world
```

### 3.4 Composite Data Types 
Besides the basic data types, Go supports other 8 composite data types, namely array, struct, pointer, function, interface, slice, map, and channel types [\[2\]][2].

#### 3.4.1 Array
Array is a list of elements of the same type. It can not have mixed-typed elements. The length of an array is its number of elements. The length of an array can be evaluated by the function ```len()```

``` go
// array.go
package main

import "fmt"

func main() {
	var arr [10]string

	arr[3] = "hello"
	fmt.Println("arr is", arr)
	fmt.Println("arr[0] is", arr[0])
	fmt.Println("arr[3] is", arr[3])
	fmt.Println("len(arr) is", len(arr))

  // this statement will cause error because arr is an array of string
	// arr[0] = 1
}
```

You should receive the following output
```
arr is ['', '', '', 'hello', '', '', '', '', '', '']
arr[0] is ''
arr[3] is 'hello'
len(arr) is 10
```

#### 3.4.2 Struct
Struct is a sequence of elements of the same or different types.

``` go
// struct.go
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
```
You should receive the following output:
``` text
{A 3.8}
A
3.8
```

#### 3.4.3 Pointer
Pointer stores the memory address of the variable. 

Type `*T` is a pointer that points to a value of type `T`. Its value without initialization is `nil`.

The `&` operator creates a pointer to the variable.

The `*` operator evaluates the value pointed by the pointer.

``` go
// pointer.go
package main

import (
	"fmt"
)

func main() {
	var p *int // initialize pointer variable
	i := 4
	p = &i
	fmt.Println("==Before==")
	fmt.Println("i:", i) // print value of i on its initialization
	fmt.Println("p:", p) // print pointer p (the memory address of i)
	*p = *p / 2          // take the value stored in p and divide it by 2
	fmt.Println("==After==")
	fmt.Println("i:", i) // print p again to see if the memory address changes
	fmt.Println("p:", p) // print value of i to see if its value changes
}
```
You should receive the following output:
``` text
==Before==
i: 4
p: 0xc000096008
==After==
i: 2
p: 0xc000096008
```

#### 3.4.4 Function 

Function type defines functions in Go

```go
// function.go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
}
```
When you run this code, you should receive an output `3`.

#### 3.4.5 Interface

#### 3.4.6 Slice
Slice is a data type that is associated with an array. It provides a way to access a block of consecutive elements in an array. It is initialized by a slice constructor, and each slice object associates with only one array. Unlike array, slice can extends its size as long as its sie belongs to the parent array.
``` go

```

#### 3.4.7 Map
Map data type in Go is similar to dictionary in python, where each key is attached with a value.
```go
// map.go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["ten"] = 10
	m["one"] = 1
	fmt.Println(m["ten"])
}
```
In this example, we map the word form of a number with its numeric value. You should receive `10` in the output.

#### 3.4.8 Channel 
Channel is a unique data type in Go. As Go is designed for concurrency, it creates Channel as a way to send and receive values of a specific data types across execution threads. Here is a toy example taken from [Go by Example: Channels][6]

```go
// channel.go
package main
import "fmt"

func main() {
    messages := make(chan string)
    go func() { messages <- "ping" }()
    msg := <-messages
    fmt.Println(msg)
}
```
The above code will output
```text
ping
```
## 3.5 Mutability
Go has both mutable and immutable objects [\[3\]][3]
### Mutable Go objects:
- arrays
- slices
- maps
- channels
- structs having multiple fields

### Immutable Go objects:
- interfaces
- booleans
- int
- float
- strings
- pointers
- structs having a single field

---

## References:
- [A tour of Go][1]
- [The Go Programming Language Report][2]
- [Go by Example: Channels][3]
- [The pros and cons of programming in Go][4]
- [The Go Programming Language Specification - Constant][5]
- [StackOverflow - Which types are mutable and immutable in the Google Go Language?][6]
- [The Go Programming Language Specification - Operators][7]

[1]: https://go.dev/tour/basics/8
[2]: https://kuree.gitbooks.io/the-go-programming-language-report/content/4/text.html
[3]: https://gobyexample.com/channels
[4]: https://www.willowtreeapps.com/craft/the-pros-and-cons-of-programming-in-go#:~:text=Go%20is%20a%20strongly%2C%20statically,catch%20entire%20classes%20of%20bugs.
[5]: https://go.dev/ref/spec#Constants
[6]: https://stackoverflow.com/questions/8018081/which-types-are-mutable-and-immutable-in-the-google-go-language#:~:text=Then%20the%20answer%20is%3A%20Integer,a%20string%20variable%20is%20mutable.
[7]: https://go.dev/ref/spec#Operators