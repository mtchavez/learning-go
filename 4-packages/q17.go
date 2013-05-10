package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
)

func main() {
	var st stack.Stack
	var token string
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading line")
	}

	for _, c := range str {
		switch {
		case c >= '0' && c <= '9':
			token = token + string(c)
		case c == ' ':
			if token != "" {
				r, _ := strconv.Atoi(token)
				st.Push(r)
			}
			token = ""
		case c == '+':
			x := st.Pop() + st.Pop()
			st.Push(x)
		case c == '*':
			x := st.Pop() * st.Pop()
			st.Push(x)
		case c == '-':
			x := st.Pop()
			y := st.Pop()
			st.Push(y - x)
		case c == '\n':
			fmt.Printf("result: %d\n", st.Pop())
		default:
			fmt.Println("Bad format!")
			return
		}
	}
}
