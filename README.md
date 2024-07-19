# Virtual File System
This project implements a virtual file system with user, folder, and file management capabilities. It provides a command-line interface for managing users, folders, and files. The project is organized into several packages, each with specific responsibilities.

### Packages
- **cmd**: Contains the main entry point for the application.
- **repo**: Contains the repository implementations for managing users, folders, and files in memory.
- **service**: Contains the service layer implementations for user, folder, and file management.
- **handler**: Contains the handlers for processing user, folder, and file commands.
- **repl**: Contains the REPL (Read-Eval-Print Loop) implementation for the CLI.
- **test**: Contains the integration tests for the project.
- **vfs**: Define domain object of user, folder, and file.

### Getting Started
#### Running the Application
To start the virtual file system REPL, run:
```bash
go run cmd/main.go
```

You will see a prompt where you can enter commands:
```bash
Virtual File System REPL
# 
```

## Available Commands
- `register [username]`: Register a new user.
- `create-folder [username] [foldername] [description]?`: Create a new folder.
- `delete-folder [username] [foldername]`: Delete a folder.
- `list-folders [username] [--sort-name|--sort-created] [asc|desc]`: List folders.
- `rename-folder [username] [foldername] [new-folder-name]`: Rename a folder.
- `create-file [username] [foldername] [filename] [description]?`: Create a new file.
- `delete-file [username] [foldername] [filename]`: Delete a file.
- `list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`: List files.