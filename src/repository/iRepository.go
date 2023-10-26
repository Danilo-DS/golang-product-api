package repository

type RepositoryType RepositoryTypes

type RepositoryTypes interface {
	ProductRepository | CategoryRepository
}

type IRepository[T any] interface {
	ListAll() ([]T, error)
	GetById(id uint64) (T, error)
	Save(entity T) (T, error)
	Update(entity T) error
	DeleteById(id uint64) error
	ExistsById(id uint64) (bool, error)
}
