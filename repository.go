package genericgorm

import "gorm.io/gorm"

// Repository ...
type Repository[T any] struct {
	orm *gorm.DB
}

// NewRepository ...
func NewRepository[T any](conn *gorm.DB) *Repository[T] {
	return &Repository[T]{
		orm: conn,
	}
}
