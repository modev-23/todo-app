# Go Todo CLI Application

A simple yet powerful command-line todo application written in Go that helps you manage your tasks efficiently.

## 🚀 Features

- Create and manage todos from the command line
- Store todos persistently (currently using CSV)
- Simple and intuitive command interface
- List all todos with their details
- Delete todos when completed
- Fast and lightweight

## 📋 Prerequisites

- Go 1.20 or higher

## 🛠️ Installation

1. Clone the repository:
```bash
git clone https://github.com/modev-23/todo-app.git
```

2. Navigate to the project directory:
```bash
cd todo-app
```

3. Build the application:
```bash
go build
```

## 💻 Usage

### Basic Commands

```bash
# Add a new todo
./todo-app add "Complete the project documentation"

# List all todos
./todo-app list

# Delete a todo by ID
./todo-app delete 1
```

## 🏗️ Project Structure

```
todo-app/
├── main.go
├── cmd/
│   └── root.go
├── internal/
│   ├── storage/
│   │   └── csv.go
│   └── models/
│       └── todo.go
├── go.mod
└── go.sum
```

## 🗄️ Data Storage

Currently, the application uses CSV files for data persistence. The todos are stored in a local CSV file with the following structure:

```csv
id,title,created_at,completed
1,Complete documentation,2024-10-31T10:00:00Z,false
```

## 🚧 Upcoming Features

The following features are planned for the next iteration:

1. **Database Migration**
   - Replace CSV storage with BoltDB for improved performance and reliability
   - Implement proper indexing for faster queries

2. **Enhanced Todo Management**
   - Inline search functionality for quick todo lookup
   - Mark todos as done without deleting them
   - Todo categories and priorities

3. **Interactive TUI**
   - Implementation of an interactive Terminal User Interface using the [bubbletea](https://github.com/charmbracelet/bubbletea) package
   - Visual task management
   - Keyboard shortcuts for common operations
   - Real-time updates and notifications

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

modev-23

## ⭐ Show your support

Give a ⭐️ if this project helped you!
