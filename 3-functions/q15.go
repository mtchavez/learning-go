package main

import (
  "fmt"
)

func plusTwo() func(int) int {
  f := func(x int) int {
    return x + 2
  }
  return f
}

func plusX(x int) func(int) int {
  f := func(y int) int {
    return x + y
  }
  return f
}

func main() {
  p := plusTwo()
  fmt.Printf("%v\n", p(2))

  q := plusX(2)
  fmt.Printf("%v\n", q(7))
}
