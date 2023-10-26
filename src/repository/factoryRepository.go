package repository

import (
	"product-api/src/config/db"
)

type FactoryMethod[T IRepository[any]] interface {
	GetRepository() (T, error)
}

type FactoryMethodImpl[T RepositoryType] struct {
}

func (f *FactoryMethodImpl[T]) GetRepository() (*T, error) {
	connection, err := db.StartConnection()
	return &T{connection}, err
}

func InitFactory[T RepositoryType]() *FactoryMethodImpl[T] {
	return &FactoryMethodImpl[T]{}
}
