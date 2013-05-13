package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var nLines int

// func cat() {

// }

func main() {
	flag.IntVar(&nLines, "n", 0, "Number of lines to print out")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please provide a file to cat")
		return
	}
	if nLines <= 0 {
		fmt.Println("Print whole file")
	}

	filename := args[0]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := r.ReadString('\n')
	i := 0
	for e == nil {
		if nLines != 0 {
			i++
		}
		if i > nLines {
			break
		}
		fmt.Println(s)
		s, e = r.ReadString('\n')
	}
}

func usage() {
	fmt.Printf("goprog [-n] file")
	flag.PrintDefaults()
	os.Exit(1)
}
