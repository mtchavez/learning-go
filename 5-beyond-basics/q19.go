package main

import (
  "fmt"
)

type Items interface { }

func Map(items []Items, mapFun func(Items) Items) []Items {
  for i, item := range items {
    items[i] = mapFun(item)
  }
  return items
}

func add2(iface Items) Items {
  switch v := iface.(type) {
    case int:
      return v * 2
    case string:
      return v + v
  }
  return iface
}

func main() {
  ilist := []Items {1, 2, 3, 4, 5, 6, 7, 8, 9}
  slist := []Items {"H", "e", "l", "l", "o", "!"}
  fmt.Printf("Integer List: %v\n", Map(ilist, add2))
  fmt.Printf("String List: %v\n", Map(slist, add2))
}
