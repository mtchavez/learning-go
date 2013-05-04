package main

import (
  "fmt"
)

func ex1() {
  fmt.Println("ex1")
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }
  fmt.Println("")
}

func ex2() {
  fmt.Println("ex3")
  i := 0
Loop:
  fmt.Println(i)
  i++
  if i < 10 {
    goto Loop
  }
  fmt.Println("")
}

func ex3() {
  fmt.Println("ex3")
  array := [10]int{}
  for i := 0; i < len(array); i++ {
    array[i] = i
  }
  fmt.Printf("%v\n", array)
}

func main() {
  ex1()
  ex2()
  ex3()
}
