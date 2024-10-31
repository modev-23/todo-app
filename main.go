package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/modev-23/todo-app/commands"
	"github.com/rodaine/table"
)

type App struct {
	todos    []*commands.Todo
	filename string
	reader   *bufio.Reader
}

func NewApp(filename string) (*App, error) {
	todos, err := commands.LoadCsvFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load initial data: %w", err)
	}

	return &App{
		todos:    todos,
		filename: filename,
		reader:   bufio.NewReader(os.Stdin),
	}, nil
}

func (app *App) Run() error {
	fmt.Println("Todo App Started - Type 'help' for commands")

	for {
		fmt.Print("> ")
		input, err := app.reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		command := strings.TrimSpace(input)
		if err := app.processCommand(command); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func (app *App) processCommand(command string) error {
	switch command {
	case "help":
		app.printHelp()

	case "list":
		app.listTodos()

	case "new":
		return app.createTodo()

	case "save":
		return app.saveTodos()

	case "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)

	case "":
		// Ignore empty input
		return nil

	default:
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
}

func (app *App) printHelp() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  list     - Show all todos")
	fmt.Println("  new      - Create a new todo")
	fmt.Println("  save     - Save changes to file")
	fmt.Println("  exit     - Exit the application")
	fmt.Println("  help     - Show this help message")
}

func (app *App) listTodos() {
	headerFmt := color.New(color.FgCyan, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgBlue).SprintfFunc()

	tbl := table.New("ID", "Status", "Importance", "Details")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, todo := range app.todos {
		tbl.AddRow(todo.Id, todo.Status.String(), todo.Importance.String(), todo.Description)
	}

	tbl.Print()
}

func (app *App) createTodo() error {
	fmt.Print("Enter todo description: ")
	description, err := app.reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read description: %w", err)
	}

	description = strings.TrimSpace(description)
	if description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	todo := commands.NewTodo(len(app.todos)+1, description)
	app.todos = append(app.todos, todo)

	fmt.Println("Todo created successfully!")
	return app.saveTodos() // Auto-save after creation
}

func (app *App) saveTodos() error {
	if err := commands.SaveCSV(app.filename, app.todos); err != nil {
		return fmt.Errorf("failed to save todos: %w", err)
	}

	fmt.Println("Todos saved successfully!")
	return nil
}

func ensureFile(filename string) error {
	// Create the file if it doesn't exist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		dir := filepath.Dir(filename)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		// Write CSV header
		_, err = file.WriteString("Id,Importance,Status,Description\n")
		if err != nil {
			return fmt.Errorf("failed to write CSV header: %w", err)
		}
	}
	return nil
}

func main() {
	filename := "todo.csv"

	if err := ensureFile(filename); err != nil {
		fmt.Printf("Failed to ensure file exists: %v\n", err)
		os.Exit(1)
	}

	app, err := NewApp(filename)
	if err != nil {
		fmt.Printf("Failed to initialize app: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(); err != nil {
		fmt.Printf("Application error: %v\n", err)
		os.Exit(1)
	}
}
