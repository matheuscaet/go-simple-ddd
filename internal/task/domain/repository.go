package domain

type TaskRepository interface {
	Save(task *Task) error
	List() ([]*Task, error)
	FindByID(id string) (*Task, error)
	Delete(id string) error
}
