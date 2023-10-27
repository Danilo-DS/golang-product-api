package repository

import . "product-api/src/models"

// Custom type
type EntityType EntityTypes

// Custom type
type RepositoryType RepositoryTypes

// Associated types
type EntityTypes interface {
	Product | Category
}

// Associated types
type RepositoryTypes interface {
	ProductRepository | CategoryRepository
}

// IRepository is an generic interface
type IRepository[T any] interface {
	ListAll() ([]T, error)
	GetById(id uint64) (T, error)
	Save(entity T) (T, error)
	Update(entity T) error
	DeleteById(id uint64) error
	ExistsById(id uint64) (bool, error)
}
