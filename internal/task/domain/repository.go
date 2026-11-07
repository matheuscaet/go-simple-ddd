package domain

import "context"

type TaskRepository interface {
	Save(ctx context.Context, task *Task) error
	List(ctx context.Context) ([]*Task, error)
	FindByID(ctx context.Context, id string) (*Task, error)
	Delete(ctx context.Context, id string) error
}
