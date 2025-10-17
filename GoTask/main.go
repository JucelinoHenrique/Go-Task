package main

import (
	"flag"
	"fmt"
	"gotask/internal/storage"
	"gotask/internal/task"
	"log"
	"os"
	"strconv"
)

var tasks []task.Task

func main() {
	loadedTasks, err := storage.LoadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	tasks = loadedTasks
	fmt.Printf("GoTask - Task Manager, Tasks: %d\n", len(tasks))
	addDescription := flag.String("add", "", "Description of new task")
	addTitle := flag.String("title", "", "Title of new task")
	listFlag := flag.Bool("list", false, "List all tasks")
	doneID := flag.String("done", "", "Mark task as completed by ID")
	flag.Parse()

	if *addDescription != "" && *addTitle != "" {
		handleAddTask(*addTitle, *addDescription)
	}

	if *doneID != "" {
		id, parseErr := strconv.Atoi(*doneID)
		if parseErr != nil {
			log.Fatalf("Invalid task ID: %v", parseErr)
		}
		handleDoneTask(tasks, id)
		return
	}

	if *listFlag {
		handleListTasks()
	}

	if len(os.Args) == 1 {
		fmt.Println("Welcome to Gotask!. Use -h for help.")
	}

}

func handleListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Tasks:")
	for _, t := range tasks {
		status := t.GetStatusString()
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", t.ID, t.Title, t.Description, status)
	}
}

func handleAddTask(title, description string) {
	newTask := task.Task{
		ID:          task.GetNextID(tasks),
		Title:       title,
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, newTask)
	err := storage.SaveTasks(tasks)
	if err != nil {
		log.Fatalf("Failed to save tasks: %v", err)
	}
	fmt.Printf("Added task: ID %d, Title: %s\n", newTask.ID, newTask.Title)
}

func handleDoneTask(tasks []task.Task, id int) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].MarkCompleted()
			err := storage.SaveTasks(tasks)
			if err != nil {
				log.Fatalf("Failed to save tasks: %v", err)
			}
			fmt.Printf("Marked task ID %d as completed.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}
