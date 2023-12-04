package main

import (
	"fmt"
	"log"
)

func main() {
	fileNames, err := parser() // Call the parser function
	if err != nil {
		log.Fatal(err)
	}

	// Print the list of file names
	fmt.Println("Files in the directory:")
	for _, name := range fileNames {
		fmt.Println(name)
	}

	rename(fileNames)
}
