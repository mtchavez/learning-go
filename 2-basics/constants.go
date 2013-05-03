package main

import "fmt"

const (
  a = iota
  b = iota
)

const (
  c = iota
  d
)

const (
  e = 0
  f string = "0"
)

func main() {
  fmt.Printf("a = %d b = %d\n", a, b)
  fmt.Printf("c = %d d = %d\n", c, d)
  fmt.Printf("e = %d f = %s\n", e, f)
}
