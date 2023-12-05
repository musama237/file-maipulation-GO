package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the path to the directory you want to work with:")
	directoryPath, _ := reader.ReadString('\n')
	directoryPath = strings.TrimSpace(directoryPath) // Remove newline character

	fmt.Println("Enter choice")
	fmt.Println("1 for list Directories")
	fmt.Println("2 for rename")
	fmt.Println("3 to move file from one dir to other")
	fmt.Println("4 to create a new file")
	fmt.Println("5 to delete a file")
	fmt.Println("6 to exit")

	var choice int
	var fileNames []string
	var err error

	fmt.Println("Enter option:")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fileNames, err = parser(directoryPath) // Call the parser function
		if err != nil {
			log.Fatal(err)
		}

		// Print the list of file names
		fmt.Println("Files in the directory:")
		for _, name := range fileNames {
			fmt.Println(name)
		}

	case 2:
		fileNames, err = parser(directoryPath) // Call the parser function
		if err != nil {
			log.Fatal(err)
		}

		// Print the list of file names
		fmt.Println("Files in the directory:")
		for _, name := range fileNames {
			fmt.Println(name)
		}
		if fileNames == nil {
			fmt.Println("Please list the directories first (Option 1)")
		} else {
			rename(fileNames)
		}

	case 3:
		fileNames, err = parser(directoryPath) // Call the parser function
		if err != nil {
			log.Fatal(err)
		}

		// Print the list of file names
		fmt.Println("Files in the directory:")
		for _, name := range fileNames {
			fmt.Println(name)
		}
		if fileNames == nil {
			fmt.Println("Please list the directories first (Option 1)")
		} else {
			move(fileNames, directoryPath)
		}

	case 4:
		new_file(directoryPath)

	case 5:
		delete(directoryPath)

	case 6:
		fmt.Println("Exiting...")
		return

	default:
		fmt.Println("Invalid option, please try again.")
	}

}
