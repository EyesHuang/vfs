package repl

import (
	"bufio"
	"fmt"
	"os"
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
		args := strings.Split(input, " ")

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
