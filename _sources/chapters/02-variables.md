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


## 2.6 Reserved Keywords

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

## References:
- [Naming Conventions for Golang Functions][1]
- [What's in a name?][2]
- [Go naming convention][3]

[1]: https://www.golangprograms.com/naming-conventions-for-golang-functions.html
[2]: https://talks.golang.org/2014/names.slide#3
[3]: https://developpaper.com/golang-naming-convention/

