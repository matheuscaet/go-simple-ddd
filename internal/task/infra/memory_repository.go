package infra

import "sync"

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

func (r *MemoryRepository) Save(record *record) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.records[record.ID] = record
}

func (r *MemoryRepository) List() []*record {
	r.mu.RLock()
	defer r.mu.RUnlock()
	records := make([]*record, 0, len(r.records))
	for _, record := range r.records {
		records = append(records, record)
	}
	return records
}

func (r *MemoryRepository) FindByID(id string) (*record, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	record, exists := r.records[id]
	return record, exists
}

func (r *MemoryRepository) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.records, id)
}
