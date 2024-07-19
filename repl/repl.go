package repl

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"vfs/handler"
)

func StartREPL(handler *handler.HandlerManager) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Virtual File System REPL")
	for {
		fmt.Print("# ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := parseInput(input)

		if len(args) == 0 {
			continue
		}

		command := args[0]
		args = args[1:]

		dispatchCommand(command, args, handler)
	}
}

func dispatchCommand(command string, args []string, handler *handler.HandlerManager) {
	switch command {
	case "register":
		handler.HandleRegister(args)
	case "create-folder":
		handler.HandleCreateFolder(args)
	case "delete-folder":
		handler.HandleDeleteFolder(args)
	case "list-folders":
		handler.HandleListFolders(args)
	case "rename-folder":
		handler.HandleRenameFolder(args)
	case "create-file":
		handler.HandleCreateFile(args)
	case "delete-file":
		handler.HandleDeleteFile(args)
	case "list-files":
		handler.HandleListFiles(args)
	default:
		fmt.Println("Unrecognized command")
	}
}

// parseInput handles inputs with double quotes and splits them into tokens
func parseInput(input string) []string {
	re := regexp.MustCompile(`"([^"]*)"|(\S+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var tokens []string
	for _, match := range matches {
		if match[1] != "" {
			tokens = append(tokens, match[1])
		} else {
			tokens = append(tokens, match[2])
		}
	}
	return tokens
}
