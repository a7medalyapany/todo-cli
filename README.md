# Todo CLI - A Command Line Todo Application

This is a basic Todo Command Line Interface (CLI) application built using Go. The project is intended for **learning purposes** to demonstrate Go fundamentals such as file handling, JSON storage, user input, and basic package structure.

## Features

- Add tasks
- List tasks
- Mark tasks as done
- Delete tasks
- Tasks persist by being saved to a JSON file
- CLI interface for user interaction

## Project Structure

```bash
todo-cli/
├── main.go          # Entry point for the application
├── task/            # Handles task-related operations (add, delete, etc.)
│   └── storage.go   # Handles file storage for the tasks
└── utils/           # Contains helper functions for the application
```

## Installation

To run this project, you'll need to have [Go](https://golang.org/doc/install) installed on your machine. Once Go is installed:

1. Clone this repository:
   ```bash
   git clone https://github.com/a7medalyapany/todo-cli.git
   ```

2. Navigate to the project directory:
   ```bash
   cd todo-cli
   ```

3. Initialize Go modules (if not already initialized):
   ```bash
   go mod init todo-cli
   ```

4. Build and run the project:
   ```bash
   go run main.go
   ```

## Usage

When running the CLI application, you'll be presented with the following options:

```
1. Add Task
2. List Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
```

- Enter the corresponding number to perform the desired action.
- Tasks will be saved to a JSON file (`tasks.json`) in the root folder.

### Example

```bash
Enter choice: 1
Enter task: Finish Go project
Task added successfully!

Enter choice: 2
ID: 1 | Task: Finish Go project | Status: Pending
```

## For Learning Purposes Only

This project was created to practice Go programming and learn about file I/O, JSON, and command-line interaction in Go. It is not intended for production use, and you are welcome to improve it, extend it, or use it to further your Go learning journey.

## Contributing

Feel free to submit issues or pull requests if you would like to suggest improvements or add new features.
