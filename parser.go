package main

import (
	"os"
)

func parser(directoryPath string) ([]string, error) {

	// Read the directory specified by the user
	files, err := os.ReadDir(*&directoryPath)
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
