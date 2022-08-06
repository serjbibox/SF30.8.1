package storage

type Storage interface {
	TaskInterface
	//GetAll() ([]Task, error)
	//GetByLabel(int) ([]Task, error)
	//GetByAuthor(int) ([]Task, error)
	//Update(int) error
	//Delete(int) error
}

type TaskInterface interface {
	Create(Task) (int, error)
	GetAll() ([]Task, error)
	GetByAuthor(interface{}) ([]Task, error)
	GetByLabel(interface{}) ([]Task, error)
}
