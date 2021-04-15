package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const version = "0.1"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	var metadataFile string
	var printVersion bool

	flag.StringVar(&metadataFile, "f", "mkdeb.yaml", "path to mkdeb.yaml var")
	flag.BoolVar(&printVersion, "v", false, "print version")

	flag.Parse()

	if printVersion {
		fmt.Println(version)
		return nil
	}

	args := flag.Args()

	if len(args) != 1 {
		return errors.New("path to output directory wasn't provided")
	}

	outputDir := args[0]

	f, err := os.Open(metadataFile)
	if err != nil {
		return err
	}

	var m Metadata
	if err := yaml.NewDecoder(f).Decode(&m); err != nil {
		return err
	}

	if err := m.Validate(); err != nil {
		return err
	}

	if err := m.Generate(outputDir); err != nil {
		return err
	}

	return nil
}
