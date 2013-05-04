package main

import (
  "fmt"
)

func ex1() {
  fmt.Println("ex1")
  a := [...]int{1,2,3,4,5}
  s1 := a[2:4]
  s2 := a[1:5]
  s3 := a[:]
  s4 := a[:4]
  s5 := s2[:]

  fmt.Printf("a = %v\n" +
    "s1 = %v\n" +
    "s2 = %v\n" +
    "s3 = %v\n" +
    "s4 = %v\n" +
    "s5 = %v\n", a, s1, s2, s3, s4, s5)
}

func ex2() {
  fmt.Println("ex2")
  s0 := []int{0,0}
  s1 := append(s0, 2)
  s2 := append(s1, 3, 5, 7)
  s3 := append(s2, s0...)
  fmt.Printf("s0 = %v\n" +
    "s1 = %v\n" +
    "s2 = %v\n" +
    "s3 = %v\n", s0, s1, s2, s3)
}

func main() {
  sl := make([]int, 10)
  fmt.Println("sl = ", sl)

  array := [10]int{}
  slice := array[0:5]

  fmt.Printf("len(slice) == %d cap(slice) == %d len(array) == %d\n", len(slice), cap(slice), len(array))

  ex1()
  ex2()
}
