package main

import (
	"flag"
	"fmt"
	"os"
)

var gitCommit string    // Git hash, set by build pipeline
var buildVersion string // human readable version, set by build pipeline

func main() {
	boolPtr := flag.Bool("v", false, "version")
	flag.Parse()
	if len(flag.Args()) >= 1 && flag.Args()[0] == "version" || *boolPtr {
		fmt.Println(buildVersion, gitCommit)
		os.Exit(0)
	}

	fmt.Println("Hello World")
}
