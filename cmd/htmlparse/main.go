package main

import (
	"fmt"
	"os"

	"github.com/rdtharri/link"
)

func main() {

	// Grab arguments
	filename := os.Args[1]

	// Read in filename
	htmlReader, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("Error opening file %s: %s", filename, err)
		os.Exit(1)
	}
	defer htmlReader.Close()

	// Grab links
	links, err := link.Parse(htmlReader)
	if err != nil {
		fmt.Errorf("Error parsing html: %s", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", links)

}
