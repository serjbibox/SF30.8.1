package storage

import (
	"errors"

	"github.com/serjbibox/SF30.8.1/pkg/models"
	"github.com/serjbibox/SF30.8.1/pkg/storage/memdb"
)

type TaskMemdb struct {
	db memdb.DB
}

func NewTaskMemDb(db memdb.DB) TaskMemdb {
	return TaskMemdb{db: db}
}

//Создание новой задачи.
func (s *TaskMemdb) Create(t models.Task) (int, error) {
	s.db = append(s.db, t)
	return s.db[len(s.db)-1].ID, nil
}

//Удаление задачи по ID.
func (s *TaskMemdb) Delete(taskid uint64) error {
	return nil
}

//Запрос всех задач.
func (s *TaskMemdb) GetAll() ([]models.Task, error) {
	return s.db, nil
}

//Обновление задачи по ID.
func (s *TaskMemdb) Update(id uint64, t models.Task) (uint64, error) {
	return uint64(t.ID), nil
}

//Запрос задачи по ID.
func (s *TaskMemdb) GetById(id uint64) (*models.Task, error) {
	if id == 0 {
		return nil, errors.New("id can not be zero")
	}
	if uint64(len(s.db)) == 0 {
		return nil, errors.New("there are no entries in the table")
	}
	if id > uint64(len(s.db)) {
		return nil, errors.New("no such entry in the table")
	}
	return &s.db[id-1], nil
}

//Запрос списка задач по автору.
//Принимает 2 типа аргумента - uint64 и string.
//Если аргумент типа uint64, выводит список задач по ID автора (tasks.author_id).
//Если аргумент типа string, выводит список задач по имени автора (users.name).
func (s *TaskMemdb) GetByAuthor(p interface{}) ([]models.Task, error) {
	return s.db, nil
}

//Запрос списка задач по метке.
//Принимает 2 типа аргумента - uint64 и string.
//Если аргумент типа uint64, выводит список задач по ID метки (labels.id).
//Если аргумент типа string, выводит список задач по имени метки (labels.name).
func (s *TaskMemdb) GetByLabel(p interface{}) ([]models.Task, error) {
	return s.db, nil
}
