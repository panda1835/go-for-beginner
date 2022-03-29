# 4. Selection Control Structure

Go supports `if`, `else`, `switch`, and `defer` for the selection control structure.

## 4.1 `if`
The condition needs not be surrounded by `()`. The execution statement must stay within `{}` [\[1\]][1].

Syntax:
```go
if condition{
    statement*
}
```
Example
```go
// if.go
package main

import (
	"fmt"
)

func absolute(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	fmt.Println(absolute(5))
	fmt.Println(absolute(0))
	fmt.Println(absolute(-5))
}

```
You should receive `5,0,5` as the output.

In Go, `if` can start with a short statement to execute before the condition [\[2\]][2].

```go
// if-short-statement.go
package main

import (
	"fmt"
)

// this function calculate |x-5|
func absolute_x_minus_5(x int) int {
	if y := x - 5; y >= 0 { // here it evaluates y = x-5 first and compare it with 0
		return y
	}
	return 5 - x
}

func main() {
	fmt.Println(absolute_x_minus_5(5))
	fmt.Println(absolute_x_minus_5(0))
	fmt.Println(absolute_x_minus_5(-5))
}
```
You should receive `0,5,10` as the output.

## 4.2 `if ... else if ... else`
`else` in Go is logically similar to other languages. It must be written right after the closing curly bracket of `if`. Moreover, `else` can used variables defined in the scope of `if` [\[3\]][3].

Syntax:
```go
if condition{
    statement*
} else if {
    statement*
} else {
    statement*
}
```

```go
// if-else.go
package main

import (
	"fmt"
)

func main() {
    x := 10
	if y := x - 5; y >= 0 { 
		fmt.Println(y)
	} else{
        fmt.Println(-y)
    }
}
```
You should receive `5` as a result of the code.

```go
// if-else-if-else.go
package main

import (
	"fmt"
)

func main() {
	y := 3
	if y >= 5 {
		fmt.Println(5)
	} else if y < 2 {
		fmt.Println(2)
	} else {
		fmt.Println(y)
	}
}
```
You should receive `3` as a result of the code.

## 4.3 `switch ... case`
`switch` in Go evaluates the cases from top to bottom and only runs the succeeded case. Unlike other languages such as Java or C that need `break` to stop the searching, Go automatically use `break` whenever the case is satisfied. Moreover, Go switch cases need not be constants and can be of any basic data types (not just integers). The condition of `switch` can be left blank, in this case it means `true`.

Syntax:
```go
switch variable {
    case case1:
        statement*
    case case2:
        statement*
    ...
    default:
        statement*
}
```

Example:
```go
// switch.go
package main

import (
	"fmt"
)

func main() {
	y := 3
	switch y {
        case 1:
            fmt.Println("one")
            fmt.Println(1)
        case 2:
            fmt.Println("two")
            fmt.Println(2)
        default:
            fmt.Println("three")
            fmt.Println(3)
        
    }
}
```
The output of the above code looks like this

```text
three 
3
```
## 4.4 `defer`
A `defer` statements defers the execution of a function until other functions are executed. The output of the `defer` function is put onto a stack, and will be removed in LIFO order when other function returns [\[5\]][5].

```go
// defer.go
package main

import (
	"fmt"
)

func main() {
	fmt.Print("My ") // this will be executed first
    defer fmt.Print("Go") // this will be defered until the next underfered function is executed. It is put on the bottom of the defer execution stack
    defer fmt.Print("is ") // this will be put on top of the stack
	fmt.Print("name ") // this will be executed second
}
```
When you execute the code, the ouput would be `My name is Go`.

## 4.5 Short Circuit

```go
// short-circuit.go
package main

import (
	"fmt"
)

func isFalse() bool {
    fmt.Println("isFalse is executed")
    return false
}

func isTrue() bool {
    fmt.Println("isTrue is executed")
    return true
}

func main() {
	if isFalse() && isTrue(){
        fmt.Println("No short circuit")
    } 
}
```
If you run the code, you will get `false` which indicates that Go uses short circuit.

## Reference:
- [A Tour of Go - If][1]
- [A Tour of Go - If with a short statement][2]
- [A Tour of Go - If and else][3]
- [A Tour of Go - Switch][4]
- [A Tour of Go - Defer][5]

[1]: https://go.dev/tour/flowcontrol/5
[2]: https://go.dev/tour/flowcontrol/6
[3]: https://go.dev/tour/flowcontrol/7
[4]: https://go.dev/tour/flowcontrol/9
[5]: https://go.dev/tour/flowcontrol/12