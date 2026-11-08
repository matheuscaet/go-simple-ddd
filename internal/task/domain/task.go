package domain

import "github.com/google/uuid"

type Task struct {
	id        string
	title     string
	completed bool
}

func NewTask(title string) *Task {
	return &Task{
		id:        uuid.NewString(),
		title:     title,
		completed: false,
	}
}

func (t *Task) Update(title string, completed bool) {
	t.title = title
	t.completed = completed
}

func (t *Task) ID() string {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Completed() bool {
	return t.completed
}

func ToTask(id string, title string, completed bool) *Task {
	return &Task{
		id:        id,
		title:     title,
		completed: completed,
	}
}
