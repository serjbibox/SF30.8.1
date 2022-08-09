package storage

type TaskStorage interface {
	Create(Task) (int, error)
	GetAll() ([]Task, error)
	GetById(taskid uint64) (*Task, error)
	GetByAuthor(interface{}) ([]Task, error)
	GetByLabel(interface{}) ([]Task, error)
	Update(uint64, Task) (uint64, error)
	Delete(taskid uint64) error
}
