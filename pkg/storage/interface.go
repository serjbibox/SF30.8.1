package storage

type Storage interface {
	TaskInterface
}

type TaskInterface interface {
	Create(Task) (int, error)
	GetAll() ([]Task, error)
	GetById(uint64) (*Task, error)
	GetByAuthor(interface{}) ([]Task, error)
	GetByLabel(interface{}) ([]Task, error)
	Update(uint64, Task) (uint64, error)
	Delete(uint64) error
}
