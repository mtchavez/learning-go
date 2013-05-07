package main

import (
  "fmt"
)

func printParams(args ...int) {
  for i, arg := range args {
    fmt.Printf("arg %d is %d\n", i+1, arg)
  }
}

func main() {
  printParams(1,2,3,4,5)
  printParams(8943,34,2323,64)
}
