package main

import (
	"flag"
	"fmt"
	"os"
)

// VERSION is the version number of github-release
const VERSION = "v0.1.0"

func main() {
	version := flag.String("version", "1.0.0", "version number for the release")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("github-release - %v\n", VERSION))
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Printf("version number is %v\n", *version)
}
