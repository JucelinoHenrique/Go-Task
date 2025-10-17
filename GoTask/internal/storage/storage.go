package storage

import (
	"encoding/json"
	"gotask/internal/task"
	"os"
)

const DataFile = "tasks.json"

func LoadTasks() ([]task.Task, error) {
	data, err := os.ReadFile(DataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}
	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(DataFile, data, 0644)
}
