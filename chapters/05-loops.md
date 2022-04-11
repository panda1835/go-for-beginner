# 5. Loops

Go only has `for` loop [\[1\]][1]. 

Syntax
```go
for <init statement>; <conditional expression>; <post statement> {
    // code block
}
```
Here are a few notes about the loop components:
- init statement: executes before each iteration
- post statement: executes after each iteration
- both init and post statement are optional and can be omitted

Example:
```go
// for-loop.go
package main

import (
	"fmt"
)

func main() {
	// calculate sum of odd numbers
	sum := 0
	for i := 1; i < 20; i = i + 2 {
		sum += i
		fmt.Println(sum)
	}
}
```

You should expect the following result when executing the above script
```text
1
4
9
16
25
36
49
64
81
100
```

## Reference:
- [A Tour of Go - For][1]

[1]: https://go.dev/tour/flowcontrol/1