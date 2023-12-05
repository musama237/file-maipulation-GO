package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func delete(directoryPath string) {

	var fileName string
	fmt.Println("Enter the name of new file ")
	fmt.Scan(&fileName)
	filePath := filepath.Join(directoryPath, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("The file alread exist in the directory")
		return
	}

	os.Remove(filePath)

}
