package commands

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

// TodoCSV is used for parsing the CSV file
type TodoCSV struct {
    Id          int    `csv:"Id"`
    Importance  string `csv:"Importance"`
    Status      string `csv:"Status"`
    Description string `csv:"Description"`
}

// Todo is your actual domain model
type Todo struct {
    Id          int
    Importance  Importance
    Status      Status
    Description string
}

type Status int

const (
    Created Status = iota
    Done
    Abandoned
)

func ParseStatus(s string) (Status, error) {
    switch s {
    case "Created":
        return Created, nil
    case "Done":
        return Done, nil
    case "Abandoned":
        return Abandoned, nil
    default:
        return Created, fmt.Errorf("invalid status: %s", s)
    }
}

func (s Status) String() string {
    return [...]string{"Created", "Done", "Abandoned"}[s]
}

type Importance int

const (
    Urgent Importance = iota
    Medium
    Secondary
    Optional
)

func ParseImportance(s string) (Importance, error) {
    switch s {
    case "Urgent":
        return Urgent, nil
    case "Medium":
        return Medium, nil
    case "Secondary":
        return Secondary, nil
    case "Optional":
        return Optional, nil
    default:
        return Medium, fmt.Errorf("invalid importance: %s", s)
    }
}

func (i Importance) String() string {
    return [...]string{"Urgent", "Medium", "Secondary", "Optional"}[i]
}

func LoadCsvFile(filePath string) ([]*Todo, error) {
    in, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer in.Close()

    var todosCSV []*TodoCSV
    if err := gocsv.UnmarshalFile(in, &todosCSV); err != nil {
        return nil, fmt.Errorf("failed to unmarshal CSV: %w", err)
    }

    todos := make([]*Todo, 0, len(todosCSV))
    for _, t := range todosCSV {
        importance, err := ParseImportance(t.Importance)
        if err != nil {
            return nil, err
        }

        status, err := ParseStatus(t.Status)
        if err != nil {
            return nil, err
        }

        todo := &Todo{
            Id:          t.Id,
            Importance:  importance,
            Status:      status,
            Description: t.Description,
        }
        todos = append(todos, todo)
    }

    return todos, nil
}

func NewTodo(id int, description string) *Todo {
    return &Todo{
        Id:          id,
        Description: description,
        Importance:  Medium,
        Status:      Created,
    }
}

func (t *Todo) MarkDone() {
    t.Status = Done
}

func (t *Todo) MarkAbandoned() {
    t.Status = Abandoned
}
