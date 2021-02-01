package main

import (
	"fmt"
	"os"
)

var GitCommit string
var BuildVersion string

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "version" {
		fmt.Println(BuildVersion, GitCommit)
		os.Exit(0)
	}

	fmt.Println("Hello World")
}
