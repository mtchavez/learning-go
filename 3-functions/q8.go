package main

import "fmt"

func main() {
  i := 0
  for ; i < 10; i++ {
    fmt.Printf("%v\n", i)
  }
  fmt.Printf("%v\n", i)
}
