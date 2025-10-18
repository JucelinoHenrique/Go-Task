package task

import (
	"testing"
)

func TestGetNextID(t *testing.T) {
	tasksEmpty := []Task{}
	nextID := GetNextID(tasksEmpty)
	if nextID != 1 {
		t.Errorf("Expected next ID to be 1 for empty task list, got %d", nextID)
	}

	taskWithData := []Task{
		{ID: 1, Title: "Task 1"},
		{ID: 2, Title: "Task 2"},
		{ID: 5, Title: "Task 5"},
	}
	nextID = GetNextID(taskWithData)
	if nextID != 6 {
		t.Errorf("Expected next ID to be 6, got %d", nextID)
	}
}
