package main

import (
	"flag"
	"fmt"
	// "os"
)

const VERSION = "1.0"

func main() {
	fmt.Println("Hello World!")

	showVersion := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *showVersion {
		fmt.Println("Version:", VERSION)
	}
}
