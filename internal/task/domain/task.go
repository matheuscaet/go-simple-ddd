package domain

type Task struct {
	id        string
	title     string
	completed bool
}

func NewTask(id string, title string) *Task {
	return &Task{
		id:        id,
		title:     title,
		completed: false,
	}
}
