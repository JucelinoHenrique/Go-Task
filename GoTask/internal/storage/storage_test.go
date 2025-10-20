package storage

import (
	"gotask/internal/task"
	"os"
	"testing"
)

func TestSaveAndLoadTasks(t *testing.T) {
	fileTemp := "test_tasks.json"
	defer os.Remove(fileTemp)

	DataFile = fileTemp

	tasks := []task.Task{
		{ID: 1, Title: "Task Save", Completed: false},
		{ID: 2, Title: "Task Load", Completed: true},
	}

	if err := SaveTasks(tasks); err != nil {
		t.Fatalf("Failed to save tasks: %v", err)
	}

	loadedTasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}
	if len(loadedTasks) != len(tasks) {
		t.Fatalf("Expected %d tasks, got %d", len(tasks), len(loadedTasks))
	}

	if loadedTasks[0].ID != tasks[0].ID || loadedTasks[0].Title != tasks[0].Title || loadedTasks[0].Completed != tasks[0].Completed {
		t.Errorf("Loaded task does not match saved task")
	}
}
