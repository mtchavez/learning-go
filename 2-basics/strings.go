package main

import "fmt"

func main() {
  s := "Hello World"
  fmt.Printf("s = %s\n", s)
  // Not valid
  // s[0] = 'c'

  c := []rune(s)
  c[0] = 'c'
  s2 := string(c)

  fmt.Printf("s2 = %s\n", s2)

  s3 := "Starting part" +
      "Ending part"
  fmt.Printf("s3 = %s\n", s3)

  s4 := `Starting part
         Ending part`
  fmt.Printf("s4 = %s\n", s4)
}
