package inmemory

import (
	"context"
	"fmt"
	"gomidka/internal/models"
	"gomidka/internal/store"
	"sync"
)

type DB struct {
	data map[int]*models.Bread

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		data: make(map[int]*models.Bread),
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Create(ctx context.Context, bread *models.Bread) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[bread.ID] = bread
	return nil
}
func (db *DB) Update(ctx context.Context, bread *models.Bread) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if db.data[bread.ID] != nil {
		db.data[bread.ID] = bread
	}
	return nil
}
func (db *DB) All(ctx context.Context) ([]*models.Bread, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	breads := make([]*models.Bread, 0, len(db.data))
	for _, bread := range db.data {
		breads = append(breads, bread)
	}

	return breads, nil
}

func (db *DB) ByID(ctx context.Context, id int) (*models.Bread, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	bread, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No bread with id %d", id)
	}

	return bread, nil
}

func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
