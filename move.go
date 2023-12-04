package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func move(fileNames []string, directoryPath string) {
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the new directory location:")
	newDirectoryPath, _ := reader.ReadString('\n')
	newDirectoryPath = strings.TrimSpace(newDirectoryPath)

	fmt.Println("Enter the name of the file you want to move:")
	var fileName string
	fmt.Scan(&fileName)

	fileMoved := false

	for _, file := range fileNames {
		if fileName == file {
			oldFilePath := filepath.Join(directoryPath, file)
			newFilePath := filepath.Join(newDirectoryPath, file)

			err := os.Rename(oldFilePath, newFilePath)
			if err != nil {
				log.Println("Error in moving the file:", err)
			} else {
				fmt.Println("File moved successfully.")
				fileMoved = true
			}
			break
		}
	}

	if !fileMoved {
		fmt.Println("File not found or not moved.")
	}
}
