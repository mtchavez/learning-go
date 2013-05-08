package stack

import (
  "testing"
  "fmt"
)

func TestPush(t *testing.T) {
  var stack Stack
  stack.Push(4)
  if stack.data[0] != 4 {
    t.Errorf("Stack should have pushed 4")
  }
}

func ExamplePush() {
  var stack Stack
  stack.Push(4)
  fmt.Printf("%d\n", stack.data[0])
  // Output:
  // 4
}

func TestPopEmptyStack(t *testing.T) {
  var stack Stack
  if stack.Pop() != 0 {
    t.Errorf("Empty stack should pop a 0")
  }
}

func ExamplePopEmptyStack() {
  var stack Stack
  fmt.Printf("%d\n", stack.Pop())
  // Output:
  // 0
}
