package task

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *Task) MarkCompleted() {
	t.Completed = true
}

func (t *Task) MarkPending() {
	t.Completed = false
}

func (t *Task) GetStatusString() string {
	if t.Completed {
		return "[x]-Completed"
	}
	return "[ ] - Pending"
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
}

func (t *Task) UpdateTitle(title string) {
	t.Title = title
}

func GetNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}
