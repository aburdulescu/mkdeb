package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

var version string

func handleError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func main() {
	var metadataFile string
	flag.StringVar(&metadataFile, "f", "mkdeb.json", "path to mkdeb.json file")
	var printVersion bool
	flag.BoolVar(&printVersion, "v", false, "print version")
	flag.Parse()
	if printVersion {
		fmt.Println(version)
		return
	}
	args := flag.Args()
	if len(args) != 1 {
		handleError(errors.New("path to output directory wasn't provided"))
	}
	outputDir := args[0]
	f, err := os.Open(metadataFile)
	if err != nil {
		handleError(err)
	}
	var p Package
	if err := json.NewDecoder(f).Decode(&p); err != nil {
		handleError(err)
	}
	if err := p.Validate(); err != nil {
		handleError(err)
	}
	if err := makeOutDir(outputDir, p); err != nil {
		handleError(err)
	}
}
