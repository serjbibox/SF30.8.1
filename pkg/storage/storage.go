package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/SF30.8.1/pkg/models"
)

type Task interface {
	Create(models.Task) (int, error)
	GetAll() ([]models.Task, error)
	GetById(taskid uint64) (*models.Task, error)
	GetByAuthor(interface{}) ([]models.Task, error)
	GetByLabel(interface{}) ([]models.Task, error)
	Update(uint64, models.Task) (uint64, error)
	Delete(taskid uint64) error
}

type Storage struct {
	Task
}

func NewStoragePostgres(db *pgxpool.Pool) *Storage {
	return &Storage{
		Task: NewTaskPostgres(db),
	}
}
