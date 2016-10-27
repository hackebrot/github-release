package main

import (
	"flag"
	"fmt"
)

func main() {
	version := flag.String("version", "1.0.0", "version number for the release")
	flag.Parse()

	fmt.Printf("version number is %v\n", *version)
}
