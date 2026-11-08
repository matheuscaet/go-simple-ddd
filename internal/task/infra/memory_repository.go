package infra

import (
	"sync"

	"github.com/matheuscaet/go-simple-ddd/internal/task/domain"
)

type record struct {
	ID        string
	Title     string
	Completed bool
}

type MemoryRepository struct {
	mu      sync.RWMutex
	records map[string]*record
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		records: make(map[string]*record),
	}
}

func (r *MemoryRepository) Save(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.records[task.ID()] = &record{
		ID:        task.ID(),
		Title:     task.Title(),
		Completed: task.Completed(),
	}
	return nil
}

func (r *MemoryRepository) List() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	records := make([]*domain.Task, 0, len(r.records))
	for _, record := range r.records {
		records = append(records, toDomainTask(record))
	}
	return records, nil
}

func (r *MemoryRepository) FindByID(id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	record, exists := r.records[id]
	if !exists {
		return nil, domain.ErrTaskNotFound
	}

	return toDomainTask(record), nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.records, id)
	return nil
}

func toDomainTask(r *record) *domain.Task {
	return domain.ToTask(r.ID, r.Title, r.Completed)
}
