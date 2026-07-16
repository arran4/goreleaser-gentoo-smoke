package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "dev"

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	configFlag := flag.String("config", "", "path to config file")
	flag.Parse()

	if *versionFlag {
		fmt.Println("goreleaser-gentoo-smoke version:", version)
		os.Exit(0)
	}

	if *configFlag != "" {
		if _, err := os.Stat(*configFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error accessing config file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Using config file: %s\n", *configFlag)
	}

	fmt.Println("goreleaser-gentoo-smoke", version)
}
