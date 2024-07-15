// cmd/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Virtual File System REPL")
	for {
		fmt.Print("# ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		command := args[0]
		args = args[1:]

		switch command {
		case "register":
			if len(args) != 1 {
				fmt.Println("Usage: register [username]")
				continue
			}
		case "create-folder":
			if len(args) < 2 || len(args) > 3 {
				fmt.Println("Usage: create-folder [username] [foldername] [description]?")
				continue
			}
		case "delete-folder":
			if len(args) != 2 {
				fmt.Println("Usage: delete-folder [username] [foldername]")
				continue
			}
		case "list-folders":
			if len(args) < 1 || len(args) > 3 {
				fmt.Println("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
				continue
			}
		case "rename-folder":
			if len(args) != 3 {
				fmt.Println("Usage: rename-folder [username] [foldername] [new-folder-name]")
				continue
			}
		case "create-file":
			if len(args) < 3 || len(args) > 4 {
				fmt.Println("Usage: create-file [username] [foldername] [filename] [description]?")
				continue
			}
		case "delete-file":
			if len(args) != 3 {
				fmt.Println("Usage: delete-file [username] [foldername] [filename]")
				continue
			}
		case "list-files":
			if len(args) != 4 {
				fmt.Println("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
				continue
			}
		default:
			fmt.Println("Unrecognized command")
		}
	}
}
