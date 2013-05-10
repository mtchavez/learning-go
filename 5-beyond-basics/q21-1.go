package main

import (
  "fmt"
  "container/list"
)


func main() {
  l := list.New()
  for _, val := range []int{1,2,4} {
    l.PushBack(val)
  }
  for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
  }
}
