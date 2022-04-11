# 6. Functions

Functions in Go can take zero or more parameters. They can be the same type or of different types. Functions can also return zero, one, or more values. [\[1\]]

Syntax
```text
func <function name> (<param 1> <type 1>, <param 2> <type 2>, ...) <return type>{
    // code block

    return <return values>
}
```

## 6.1 Example
### Normal function
```go
// multiply.go

package main

import (
	"fmt"
)

func multiply(a float32, b float32) float32 {
	return a * b
}

func main() {
	fmt.Println(multiply(12, 20))
}
```
Output
```text
240
```
The above function takes in 2 float numbers and returns their product. 

### Function with no parameter
```go
// function-no-param.go
package main

import (
	"fmt"
)

func take_nothing() {
	fmt.Println("This function can run without any parameter")
}

func main() {
	take_nothing()
}
```
Output
```text
This function can run without any parameter
```

### Function with multiple returns
```go
// swap.go

package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var m, n = 10, 20
	fmt.Println("Before swap m, n =", m, n)

	m, n = swap(m, n)

	fmt.Println("After swap m, n =", m, n)
}
```

Output
```text
Before swap m, n = 10 20
After swap m, n = 20 10
```

## Recursive function
```go
// factorial.go
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return n * factorial(n-1)
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
}
```

Output
```text
1! = 1
2! = 2
3! = 6
4! = 24
5! = 120
6! = 720
7! = 5040
8! = 40320
9! = 362880
10! = 3628800
```

## 6.2 "Naked" return
Return in Go can be named, and that name is used implicitly in the return statement. The function with named returns does not specify return value after `return` command. [\[2\]][2]

Syntax
```text
func <function name> (<param 1> <type 1>, <param 2> <type 2>, ...) (<return variable> <return type>){
    // code block

    return
}
```

Example
```go
// split-string.go
package main

import (
	"fmt"
)

func split_string(a string, i int) (a1, a2 string) {
	// this function splits a string at a specific index
	a1, a2 = a[:i], a[i:]
	return
}

func main() {
	var s = "Hello World!"
	var i = 4
	var s1, s2 = split_string(s, i)
	fmt.Printf("String '%s' is divided at %d into '%s' and '%s'", s, i, s1, s2)
}
```

Output
```text
String 'Hello World!' is divided at 4 into 'Hell' and 'o World!'
```

## 6.3 Pass-by-value or Pass-by-reference

Go is a pass-by-value language. You can see it clearly in the modified swap code.
```go
// modified-swap.go

package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var m, n = 10, 20
	fmt.Println("Before swap m, n =", m, n)

	swap(m, n)

	fmt.Println("After swap m, n =", m, n)
}
```
Output
```
Before swap m, n = 10 20
After swap m, n = 10 20
```

## Reference:
- [A Tour of Go - Functions][1]
- [A Tour of Go - Named return values][2]

[1]: https://go.dev/tour/basics/4
[2]: https://go.dev/tour/basics/7