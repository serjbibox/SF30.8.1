package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/SF30.8.1/pkg/models"
)

const (
	ID_TYPE = 1 + iota
	NAME_TYPE
)

type TaskPostgres struct {
	db *pgxpool.Pool
}

func NewTaskPostgres(db *pgxpool.Pool) *TaskPostgres {
	return &TaskPostgres{db: db}
}

//Создание новой задачи.
func (s *TaskPostgres) Create(t models.Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content, author_id)
		VALUES ($1, $2, $3) RETURNING id;
		`,
		t.Title,
		t.Content,
		t.AuthorID,
	).Scan(&id)
	return id, err
}

//Удаление задачи по ID.
func (s *TaskPostgres) Delete(taskid uint64) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks 
		WHERE id = $1	
		`,
		taskid,
	)
	return err
}

//Запрос всех задач.
func (s *TaskPostgres) GetAll() ([]models.Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT *
		FROM tasks
		ORDER BY id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

//Обновление задачи по ID.
func (s *TaskPostgres) Update(id uint64, t models.Task) (uint64, error) {

	err := s.db.QueryRow(context.Background(), `
		UPDATE tasks 
		SET opened = $1, 
			closed = $2, 
			author_id = $3, 
			assigned_id = $4, 
			title = $5, 
			content = $6
		WHERE id = $7	
		RETURNING id;
		`,
		t.Opened,
		t.Closed,
		t.AuthorID,
		t.AssignedID,
		t.Title,
		t.Content,
		id,
	).Scan(&id)
	return id, err
}

//Запрос задачи по ID.
func (s *TaskPostgres) GetById(id uint64) (*models.Task, error) {
	var err error
	var t models.Task
	err = s.db.QueryRow(context.Background(), `
			SELECT * 
			FROM tasks
			WHERE tasks.id = $1
			ORDER BY id;
		`,
		id,
	).Scan(
		&t.ID,
		&t.Opened,
		&t.Closed,
		&t.AuthorID,
		&t.AssignedID,
		&t.Title,
		&t.Content,
	)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

//Запрос списка задач по автору.
//Принимает 2 типа аргумента - uint64 и string.
//Если аргумент типа uint64, выводит список задач по ID автора (tasks.author_id).
//Если аргумент типа string, выводит список задач по имени автора (users.name).
func (s *TaskPostgres) GetByAuthor(p interface{}) ([]models.Task, error) {
	var rows pgx.Rows
	var err error
	switch p := p.(type) {
	case uint64:
		rows, err = s.db.Query(context.Background(),
			buildAuthorQuery(ID_TYPE),
			p,
		)
		if err != nil {
			return nil, err
		}
	case string:
		rows, err = s.db.Query(context.Background(),
			buildAuthorQuery(NAME_TYPE),
			p,
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid type of query parameter")
	}
	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

//Запрос списка задач по метке.
//Принимает 2 типа аргумента - uint64 и string.
//Если аргумент типа uint64, выводит список задач по ID метки (labels.id).
//Если аргумент типа string, выводит список задач по имени метки (labels.name).
func (s *TaskPostgres) GetByLabel(p interface{}) ([]models.Task, error) {
	var rows pgx.Rows
	var err error
	switch p := p.(type) {
	case uint64:
		rows, err = s.db.Query(context.Background(),
			buildLabelQuery(ID_TYPE),
			p,
		)
		if err != nil {
			return nil, err
		}
	case string:
		rows, err = s.db.Query(context.Background(),
			buildLabelQuery(NAME_TYPE),
			p,
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid type of query parameter")
	}
	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}

func buildLabelQuery(t int) string {
	switch t {
	case ID_TYPE:
		return `
		SELECT 
			tasks.id, 
			tasks.opened, 
			tasks.closed, 
			tasks.author_id,
			tasks.assigned_id, 
			tasks.title, 
			tasks.content
		FROM tasks, tasks_labels, labels
		WHERE labels.id = $1
		AND tasks_labels.task_id = tasks.id
		AND tasks_labels.label_id = labels.id
		ORDER BY tasks.id;
		`
	case NAME_TYPE:
		return `
		SELECT 
			tasks.id, 
			tasks.opened, 
			tasks.closed, 
			tasks.author_id,
			tasks.assigned_id, 
			tasks.title, 
			tasks.content
		FROM tasks, tasks_labels, labels
		WHERE labels.name = $1
		AND tasks_labels.task_id = tasks.id
		AND tasks_labels.label_id = labels.id
		ORDER BY tasks.id;
		`
	}
	return ""
}

func buildAuthorQuery(t int) string {
	switch t {
	case ID_TYPE:
		return `
		SELECT 
			tasks.id,
			tasks.opened,
			tasks.closed,
			tasks.author_id,
			tasks.assigned_id,
			tasks.title,
			tasks.content
		FROM tasks	
		WHERE tasks.author_id = $1 
		ORDER BY id;
		`
	case NAME_TYPE:
		return `
		SELECT 
			tasks.id,
			tasks.opened,
			tasks.closed,
			tasks.author_id,
			tasks.assigned_id,
			tasks.title,
			tasks.content
		FROM tasks, users	
		WHERE users.name = $1 AND users.id = tasks.author_id
		AND users.id = tasks.author_id
		ORDER BY id;
		`
	}
	return ""
}
