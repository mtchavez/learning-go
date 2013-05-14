package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dnssec := flag.Bool("dnssec", false, "Request DNSSEC records")
	port := flag.String("port", "53", "Set the query port")

	flag.Usage = func() {
		fmt.Printf("Usage %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *dnssec {
		fmt.Printf("Has dnssec: %v\n", *dnssec)
	}
	if *port != "" {
		fmt.Printf("Has port: %v\n", *port)
	}
}
