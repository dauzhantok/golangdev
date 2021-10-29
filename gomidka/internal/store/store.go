package store

import (
	"context"
	"gomidka/internal/models"
)

type Store interface {
	Create(ctx context.Context, bread *models.Bread) error
	All(ctx context.Context) ([]*models.Bread, error)
	ByID(ctx context.Context, id int) (*models.Bread, error)
	Update(ctx context.Context, bread *models.Bread) error
	Delete(ctx context.Context, id int) error

	// Laptops() LaptopsRepository
	// Phones() PhonesRepository
}

// electronics
//   laptops
//   phones

// TODO дома почитать, вернемся в будущих лекциях
// type LaptopsRepository interface {
// 	Create(ctx context.Context, laptop *models.Laptop) error
// 	All(ctx context.Context) ([]*models.Laptop, error)
// 	ByID(ctx context.Context, id int) (*models.Laptop, error)
// 	Update(ctx context.Context, laptop *models.Laptop) error
// 	Delete(ctx context.Context, id int) error
// }
