package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const version = "0.1"

func handleError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func main() {
	var metadataFile string
	flag.StringVar(&metadataFile, "f", "mkdeb.yaml", "path to mkdeb.yaml file")
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
	var m Metadata
	if err := yaml.NewDecoder(f).Decode(&m); err != nil {
		handleError(err)
	}
	if err := m.Validate(); err != nil {
		handleError(err)
	}
	if err := m.Generate(outputDir); err != nil {
		handleError(err)
	}
}
