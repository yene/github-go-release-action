package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

//go:embed commit.txt
var gitCommit string // Git hash, set by build pipeline
//go:embed version.txt
var buildVersion string // human readable version, set by build pipeline

func main() {
	boolPtr := flag.Bool("v", false, "version")
	flag.Parse()
	if len(flag.Args()) >= 1 && flag.Args()[0] == "version" || *boolPtr {
		fmt.Println(strings.TrimSpace(buildVersion), strings.TrimSpace(gitCommit))
		os.Exit(0)
	}

	fmt.Println("Hello World")
}
