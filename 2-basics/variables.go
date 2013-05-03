package main

import "fmt"

func var1() {
  var a int
  var b bool
  a = 15
  b = false
  fmt.Printf("var1 a = %d b = %v\n", a, b)
}

func var2() {
  a := 15
  b := false
  fmt.Printf("var2 a = %d b = %v\n", a, b)
}

func var3() {
  var (
    x int
    b bool
  )
  x = 0
  b = true
  fmt.Printf("var3 x = %d b = %v\n", x, b)
}

func var4() {
  a, b := 20, 16
  _, d := 34, 35
  fmt.Printf("var4 a = %d b = %d d = %d\n", a, b, d)
}

func main() {
  var1()
  var2()
  var3()
  var4()
  // Invalid since it is not used
  // and compiler raises warnings for
  // unused variables
  // var i int
}
