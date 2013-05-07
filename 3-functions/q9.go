package main

import (
"fmt"
)

type Stack struct {
  next_index int
  data [10]int
}

func (s *Stack) Push(val int) {
  if s.next_index + 1 > 9 {
    return
  }
  s.data[s.next_index] = val
  s.next_index += 1
}

func (s *Stack) Pop() int {
  if s.next_index == 0 {
    return 0
  }
  s.next_index -= 1
  return s.data[s.next_index]
}

func main() {
  var stack Stack
  stack.Push(42)
  stack.Push(18)
  fmt.Printf("Pop: %d\n", stack.Pop())
  fmt.Printf("Stack: %v\n", stack)

  for i := 0; i < 5; i++ {
    fmt.Printf("Pop: %d\n", stack.Pop())
  }
}
