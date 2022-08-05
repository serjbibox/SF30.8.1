package storage

type Storage interface {
	Create(Task) (int, error)
	GetAll() ([]Task, error)
	GetByLabel(int) ([]Task, error)
	GetByAuthor(int) ([]Task, error)
	Update(int) error
	Delete(int) error
}
