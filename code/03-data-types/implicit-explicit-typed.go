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