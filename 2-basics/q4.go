package main

import (
  "fmt"
  "unicode/utf8"
)

func q1() {
  fmt.Println("Q1")
  str := "A"
  for i := 0; i < 100; i++ {
    fmt.Println(str)
    i++
    str += "A"
  }
  fmt.Println()
}

func q2() {
  fmt.Println("Q2")
  str := "asSASA ddd dsjkdsjs dk"
  chars := []rune(str)
  dict := map[string]int {}
  for i := 0; i < len(chars); i++ {
    key := string(chars[i])
    _, ok := dict[key]
    if ok {
      dict[key] += 1
    } else {
      dict[key] = 1
    }
  }
  fmt.Println(dict)
  fmt.Printf("'%s' takes up %d bytes\n", str, utf8.RuneCount([]byte(str)))
  fmt.Println()
}

func q4() {
  fmt.Println("Q4")
  str := "foobar"
  bytes := []byte(str)
  for i, j := 0, len(bytes) - 1; i < j; i, j = i + 1, j - 1 {
    bytes[i], bytes[j] = bytes[j], bytes[i]
  }
  fmt.Printf("%s reversed is %s\n", str, string(bytes))
  fmt.Println()
}

func main() {
  q1()
  q2()
  q4()
}
