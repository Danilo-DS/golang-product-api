package repository

import (
	"product-api/src/config/db"
)

// FactoryMethod is an interface  type IRepository
type FactoryMethod[T IRepository[any]] interface {
	GetRepository() (T, error)
}

// FactoryMethodImpl is an implementation FactoryMethod
type FactoryMethodImpl[T RepositoryType] struct {
}

// GetRepository return a repository by specified type
func (f *FactoryMethodImpl[T]) GetRepository() (*T, error) {
	connection, err := db.StartConnection()
	return &T{connection}, err
}

// InitFactory initialize repository factory
func InitFactory[T RepositoryType]() *FactoryMethodImpl[T] {
	return &FactoryMethodImpl[T]{}
}
