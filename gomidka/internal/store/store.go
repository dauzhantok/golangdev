package store

import (
	"context"
	"gomidka/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	Breads() BreadsRepository
	Categories() CategoriesRepository
}
type BreadsRepository interface {
	Create(ctx context.Context, bread *models.Bread) error
	All(ctx context.Context) ([]*models.Bread, error)
	ByID(ctx context.Context, id int) (*models.Bread, error)
	Update(ctx context.Context, bread *models.Bread) error
	Delete(ctx context.Context, id int) error
}
type CategoriesRepository interface {
	Create(ctx context.Context, bread *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
	ByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, bread *models.Category) error
	Delete(ctx context.Context, id int) error
}