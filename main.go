package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func handleError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func main() {
	var metadataFile string
	var outputDir string
	flag.StringVar(&metadataFile, "f", "mkdeb.json", "path to metadata file")
	flag.StringVar(&outputDir, "o", "mkdeb.out", "path to output directory")
	flag.Parse()
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
