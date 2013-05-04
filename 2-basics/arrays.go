package main

import (
  "fmt"
)


func main() {
  var arr [10]int
  arr[0] = 42
  arr[1] = 13
  fmt.Printf("The first elem is %d\n", arr[0])

  a := [...]int{1,2,3}
  fmt.Println("a = ", a)

  b :=[2][2]int{ [2]int{1,2}, [2]int{3,4} }
  fmt.Println("b = ", b)

  c := [2][2]int{ {1,2}, {3,4} }
  fmt.Println("c = ", c)
}
