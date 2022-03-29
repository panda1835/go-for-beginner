# 2. Variables

## 2.1 Variable Declaration 

### 2.1.2 The Full Format
Syntax: 
``` go
var variableList type 
```

For example:
``` go
var age, height, weight int
```
### 2.2.2 Other Ways to Declare Variables:
- Variables with initializers (if the initializers are present, type can be omitted and the variable will take the type of the initializers)
  
  Syntax: 
  ``` go
  var variableList type = initializers
  var variableList = initializers
  ```

  For example:
  ``` go
  var age, height, weight int = 18, 168, 56
  var age, height, weight = 18, 168, 56
  ```
- Inside the function, variables can be declared implicitly with shortcut ```:=``` (This is not applicable for outside of function)

  Syntax:
  ``` go
  func function(){
    variableName := value
  }
  ```
  For example:
  ``` go
  package main

  import "fmt"

  func main() {
    a := true
    fmt.Println(a)
  }
  ```

## 2.3 Naming
### 2.3.1 Naming Criteria
The naming in Go must follow all the criteria [\[1\]][1]:
  1. A name must begin with a letter, and can have any number of additional letters and numbers.
  2. A function name cannot start with a number.
  3. A function name cannot contain spaces.
  4. If the functions/variables with names that start with an uppercase letter will be exported to other packages. If the function/variable name starts with a lowercase letter, it won't be exported to other packages, but you can call this function/variable within the same package (*exported* means that function/variable can be called outside of the package). For example, ```Pi``` is an exported variable from package ```math```, while ```pi``` is not. So, we can call ```Pi``` from ```math``` in our main program (sample code in file ```exported-names.go```).
      ```go
      import (
        "fmt"
        "math"
      )

      func main() {
          fmt.Println(math.Pi) // will work
          fmt.Println(math.pi) // will not work
      }
      ```
  5. If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
  6. Function names are case-sensitive (go, Go and GO are three different variables).

### 2.3.2 Naming Convention
Besides rigid rules above, Go developers also develop their own culture for naming variables [\[2\]][2]

As a rule of thumb:
  > "The greater the distance between a name's declaration and its uses, the longer the name should be." -- Andrew Gerrand

- Name should use camel case (e.g., lastName) - and strictly avoid using underscore (e.g., ~~last_name~~)
- Acronyms should be all capitalized (e.g., ServeHTTP)
- Keep the name very short and concise. For common name such as ```index```, it should be named ```i``` instead.

## 2.4 Data Types
### 2.4.1 Basic Data Types
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

Example (see ```basic-data-types.go```):
``` go
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
### 2.4.2 Composite Data Types 
Besides the basic data types, Go supports other 8 composite data types, namely array, struct, pointer, function, interface, slice, map, and channel types [\[5\]][5].

#### 2.4.2.1 Array
Array is a list of elements of the same type. It can not have mixed-typed elements. The length of an array is its number of elements. The length of an array can be evaluated by the function ```len()```

Example (see ```array.go```)
``` go
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

#### 2.4.2.2 Struct
Struct is a sequence of elements of the same or different types.

Example (see ```struct.go```)
``` go
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

#### 2.4.2.3 Pointer
Pointer stores the memory address of the variable. 

Type `*T` is a pointer that points to a value of type `T`. Its value without initialization is `nil`.

The `&` operator creates a pointer to the variable.

The `*` operator evaluates the value pointed by the pointer.

``` go
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

#### 2.4.2.4 Function 

Function type defines functions in Go

```go
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

#### 2.4.2.5 Interface

#### 2.4.2.6 Slice
Slice is a data type that is associated with an array. It provides a way to access a block of consecutive elements in an array. It is initialized by a slice constructor, and each slice object associates with only one array. Unlike array, slice can extends its size as long as its sie belongs to the parent array.
``` go

```

#### 2.4.2.7 Map
Map data type in Go is similar to dictionary in python, where each key is attached with a value.
```go
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

#### 2.4.2.8 Channel 
Channel is a unique data type in Go. As Go is designed for concurrency, it creates Channel as a way to send and receive values of a specific data types across execution threads. Here is a toy example taken from [Go by Example: Channels][6]

```go
package main
import "fmt"

func main() {
    messages := make(chan string)
    go func() { messages <- "ping" }()
    msg := <-messages
    fmt.Println(msg)
}
```

---

There are 25 keywords in Go [\[3\]][3]
```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

and 37 reserved words
```go
Constants:    true  false  iota  nil
Types:        int  int8  int16  int32  int64
			        uint  uint8  uint16  uint32  uint64  uintptr
			        float32  float64  complex128  complex64
			        bool  byte  rune  string  error
Functions:    make  len  cap  new  append  copy  close  delete
			        complex  real  imag
			        panic  recover
```


In your example code, illustrate the important data type and operations features in your language. Write code that experiments with different operations applied on variables of the same data type and operations with variables of two different types: e.g. can you add ints and floats? Is the resulting variable an int (narrowing conversion) or a float (widening conversion)?  What about division? Can you put different data types in the same array or list?  Can one data type be converted to another either implicitly or explicitly (int to float, string to int, etc)? 


Discussion questions:

Does your language have keywords or reserved words? How many?
What are the naming requirements for variables in your language?
What about naming conventions?  Are those enforced by the compiler/interpreter, or just standards in the community?

Is your language statically or dynamically typed?
Strongly typed or weakly typed?
Explicitly typed or implicitly typed? 
Are some variables mutable while others are immutable? 
What are the operators available for each data type?
Are mixed type operations allowed? If so, how are they accommodated?
At what point are identifier names and operator symbols bound in your language? For example if you declare a (variable, class name, function name), when is it bound to the type, address? When are operators (+,*, etc.) bound to their operations?
 

CODING EXAMPLE demonstrating the features above

EXAMPLE OF (one) ILLUSTRATIVE EXAMPLE:

If you put this line (or something similar) in a program and try to print x, what does it do? 

x = "5" + 6

If it doesn't compile, why?  Is there something you can do to make it compile?

 

Describe the limitations (or lack thereof) of your programming language as they relate to the coding example portion of the assignment (adding ints and floats, storing different types in lists, converting between data types). Are there other restrictions or pitfalls that the documentation mentions that you need to be aware of?
Are there built-in complex data types that are commonly used in your language? (hint: they’d probably appear fairly early in the documentation if so)


## References:
- [Naming Conventions for Golang Functions][1]
- [What's in a name?][2]
- [Go naming convention][3]
- [A tour of Go][4]
- [The Go Programming Language Report][5]
- [Go by Example: Channels][6]

[1]: https://www.golangprograms.com/naming-conventions-for-golang-functions.html
[2]: https://talks.golang.org/2014/names.slide#3
[3]: https://developpaper.com/golang-naming-convention/
[4]: https://go.dev/tour/basics/8
[5]: https://kuree.gitbooks.io/the-go-programming-language-report/content/4/text.html
[6]: https://gobyexample.com/channels


# note: xem lại phần naming convention