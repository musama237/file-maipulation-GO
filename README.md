
### Project Overview

This project is a comprehensive file manipulation tool written in Go, which operates through a command-line interface (CLI). The tool provides several functionalities that allow users to perform common file operations within a specified directory.

### Key Functionalities

1. **List Directories**: Lists all files in a specified directory.
2. **Rename Files**: Renames files in a specified directory based on user input or a defined pattern.
3. **Move File**: [Not yet implemented] Moves a file from one directory to another.
4. **Create a New File**: [Not yet implemented] Creates a new file in a specified directory.
5. **Delete a File**: [Not yet implemented] Deletes a specific file from a directory.
6. **Exit**: Exits the application.

### Technical Details

- **Language**: The tool is written in Go, leveraging its strong support for file I/O operations and efficient handling of system-level tasks.
- **Command-Line Parsing**: It uses Go's `flag` package to parse command-line arguments.
- **Switch Case Logic**: The tool utilizes a `switch` statement to handle different user choices, making the code more organized and readable.
- **Error Handling**: Proper error handling is implemented to manage issues like file not found, permission errors, etc.



Here's a summary of each function and its role within the project:

1. **`parser(directoryPath string)`:** This function lists all the files in a given directory. It returns a slice of file names (`[]string`) and any error encountered during the listing.

2. **`rename(fileNames []string)`:** This function prompts the user for a file name from the listed files and a new name for that file. If the file exists, it is renamed to the new name using `os.Rename`.

3. **`move(fileNames []string, directoryPath string)`:** After listing files in a given directory, this function allows the user to move a selected file to a new directory. It checks for file existence in the current directory and moves the file to the user-specified location.

4. **`new_file(directoryPath string)`:** This function prompts the user to enter a name for a new file. It checks if the file already exists in the specified directory and, if not, creates the new file in that directory.

5. **`delete(directoryPath string)`:** This function asks the user for a file name and deletes the file if it exists in the provided directory. It performs a check before deletion to ensure the file is present.

6. **`main()`:** This is the entry point of the CLI tool. It prompts the user to enter a directory path to work with and provides a menu with options to list, rename, move, create, and delete files, as well as an option to exit the program. Depending on the user's choice, it calls the corresponding function and handles any errors.

The CLI tool is designed to be user-friendly, providing clear prompts and instructions. It leverages Go's `os` and `path/filepath` packages for file operations and `bufio` and `fmt` for input/output interactions.

The tool is interactive, with a `switch` statement in the `main` function that directs the flow of execution based on the user's menu selection. It is robust, with error handling present to log issues and prevent crashes during operations like file reading, renaming, and deletion.

This project illustrates the application of Go's standard library to build a practical tool for file management. It showcases the language's capabilities in handling file systems and user input, making it suitable for tasks like scripting and automation in system administration or development environments.