package task

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func SaveToFile() error {
	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadFromFile() error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil // File doesn't exist, skip loading
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		return err
	}

	for _, task := range Tasks {
		if task.ID > idCounter {
			idCounter = task.ID
		}
	}
	return nil
}
