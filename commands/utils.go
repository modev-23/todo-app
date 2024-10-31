package commands

import (
	"encoding/csv"
	"os"
	"strconv"
)

func SaveCSV(filename string, todos []*Todo) error {
	// Open the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write the header row
	err = writer.Write([]string{"Id", "Importance", "Status", "Description"})
	if err != nil {
		return err
	}

	// Write the updated data
	for _, todo := range todos {
		row := []string{
			strconv.Itoa(todo.Id),
			todo.Importance.String(),
			todo.Status.String(),
			todo.Description,
		}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	// Flush the writer
	writer.Flush()
	return writer.Error()
}
