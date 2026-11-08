package application

import "github.com/matheuscaet/go-simple-ddd/internal/task/domain"

type TaskInput struct {
	Title string `json:"title"`
}

type TaskOutput struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func toOutput(task *domain.Task) *TaskOutput {
	return &TaskOutput{
		ID:        task.ID(),
		Title:     task.Title(),
		Completed: task.Completed(),
	}
}

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(input *TaskInput) (*TaskOutput, error) {
	task := domain.NewTask(input.Title)
	if err := s.repo.Save(task); err != nil {
		return nil, err
	}
	return toOutput(task), nil
}

func (s *TaskService) List() ([]*TaskOutput, error) {
	tasks, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	outputs := make([]*TaskOutput, len(tasks))
	for i, task := range tasks {
		outputs[i] = toOutput(task)
	}
	return outputs, nil
}

func (s *TaskService) FindByID(id string) (*TaskOutput, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return toOutput(task), nil
}

func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}
