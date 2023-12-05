package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func new_file(directoryPath string) {

	var fileName string
	fmt.Println("Enter the name of new file ")
	fmt.Scan(&fileName)
	filePath := filepath.Join(directoryPath, fileName)

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		fmt.Println("The file alread exist in the directory")
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error in creating file", err)
		return

	}
	defer file.Close()
}
