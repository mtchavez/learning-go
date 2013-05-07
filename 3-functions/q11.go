package main

import (
  "fmt"
)

func fibonacci(num int) (sequence []int) {
  if num <= 0 {
    num = 0
  }
  sequence = make([]int, num)
  if num == 0 {
    return
  }
  sequence[0] = 1
  if num == 1 {
    return
  }
  sequence[1] = 1
  for i := 2; i < num; i++ {
    sequence[i] = sequence[i-1] + sequence[i-2]
  }
  return
}

func main() {
  fmt.Printf("Fibonnaci -2 numbers in %v\n", fibonacci(-2))
  fmt.Printf("Fibonnaci 0 numbers in %v\n", fibonacci(0))
  fmt.Printf("Fibonnaci 1 number in %v\n", fibonacci(1))
  fmt.Printf("Fibonnaci 10 numbers in %v\n", fibonacci(10))
  fmt.Printf("Fibonnaci 20 numbers in %v\n", fibonacci(20))
}
