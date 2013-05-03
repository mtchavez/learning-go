package main

import (
  "fmt"
  "os"
)

func fun1(x int) bool {
  if x > 0 {
    return true
  }
  return false
}

func true_fun() {
  if true && true {
    fmt.Println("true && true = true")
  }

  if !false{
    fmt.Println("!false = true")
  }
}

func open_file(name string) {
  _, err := os.Open(name)
  if err != nil {
    fmt.Println("Error opening", name)
    return
  }
  fmt.Println("Found file", name)
}

func goto_fun() {
  i := 0
Here:
  fmt.Println(i)
  i++
  if i < 10 {
    goto Here
  }
}

func sum() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum++
  }
  fmt.Println("sum =", sum)
}

func reverse(a[] int) {
  fmt.Println("before reverse", a)
  for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
    a[i], a[j] = a[j], a[i]
  }
  fmt.Println("after reverse", a)
}

func breaking() {
  fmt.Println("Breaking")
J:  for j := 0; j < 5; j++ {
      for i := 0; i < 10; i++ {
          if i > 5 {
            break J
          }
        fmt.Println(i)
      }
    }
}

func ranges() {
  fmt.Println("Ranges")
  list := []string {"a", "b", "c", "d", "e", "f"}

  for k, v := range list {
    fmt.Printf("%d => %s\t", k, v)
  }
  fmt.Println()

  fmt.Println("String Range")
  for pos, char := range "aΦx" {
    fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
  }
}

func unhex(c byte) byte {
  switch {
  case '0' <= c && c <= '9':
    return c - '0'
  case 'a' <= c && c<= 'f':
    return c - 'a' + 10
  case 'A' <= c && c <= 'F':
    return c - 'A' + 10
  }
  return 0
}

func switches(i int) {
  switch i {
  case 0:
    fmt.Println("got a 0")
  case 1:
    fmt.Println("got a 1")
  default:
    fmt.Println("not 1 or 0")
  }
}

func shouldEscape(c byte) bool {
  switch c {
  case ' ', '?', '&', '=', '#', '+':
    return true
  }
  return false
}

func Compare(a string, b string) int {
  for i := 0; i < len(a) && i < len(b); i++ {
    switch {
    case a[i] > b[i]:
      return 1
    case a[i] < b[i]:
      return -1
    }
  }
  switch {
  case len(a) > len(b):
    return 1
  case len(a) < len(b):
    return -1
  }
  return 0
}

func main() {
  fmt.Println("fun1(4) =", fun1(4))
  fmt.Println("fun1(-1) =", fun1(-1))
  true_fun()
  open_file("random_file.txt")
  goto_fun()
  sum()
  reverse([]int{1,2,3,4,5,6,7,8,9})
  breaking()
  ranges()
  fmt.Println("unhex(a)=", unhex('a'))
  fmt.Println("unhex(0)=", unhex('0'))
  fmt.Println("unhex(Z)=", unhex('Z'))
  switches(1)
  switches(0)
  switches(3)
  fmt.Println("escape(&)=", shouldEscape('&'))
  fmt.Println("escape(P)=", shouldEscape('P'))
  fmt.Println("(haile == selassie) =>", Compare("haile", "selassie"))
  fmt.Println("(apple == apple) =>", Compare("apple", "apple"))
}
