package main

import (
	"flag"
	"fmt"
	"os"
)

func parser() ([]string, error) {
	// Define command-line flags
	directory := flag.String("dir", ".", "Look in the directory")
	pattern := flag.String("pattern", "", "Pattern for renaming")

	// Parse the flags
	flag.Parse()

	// Print the values of the flags
	fmt.Println("Directory:", *directory)
	fmt.Println("Pattern:", *pattern)

	// Read the directory specified by the user
	files, err := os.ReadDir(*directory)
	if err != nil {
		return nil, err // Return an error if unable to read the directory
	}

	var fileNames []string // Initialize a slice to hold the file names

	// Iterate over the files in the directory
	for _, file := range files {
		fileNames = append(fileNames, file.Name()) // Add the file name to the slice
	}

	return fileNames, nil // Return the slice of file names
}
