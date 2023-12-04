package main

import (
	"fmt"
	"log"
	"os"
)

func rename(fileNames []string) {
	var fileTorename, newname string
	fmt.Println("Enter the file name that you want to rename : ")
	fmt.Scan(&fileTorename)

	for _, file := range fileNames {
		if file == fileTorename {
			fmt.Println("Enter new name : ")
			fmt.Scan(&newname)
			err := os.Rename(fileTorename, newname)
			if err != nil {
				log.Printf("Error in renaming the file ")
				return
			}

			fmt.Printf("The file %s has been renamed to %s", fileTorename, newname)
			break
		}

	}

}
