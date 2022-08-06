package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/SF30.8.1/pkg/storage"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

//Создание новой задачи.
func (s *Storage) Create(t storage.Task) (int, error) {
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

//Обновление задачи по ID.
func (s *Storage) Update(id uint64, t storage.Task) (uint64, error) {

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

//Удаление задачи по ID.
func (s *Storage) Delete(id uint64) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks 
		WHERE id = $1	
		`,
		id,
	)
	return err
}

//Запрос всех задач.
func (s *Storage) GetAll() ([]storage.Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT *
		FROM tasks
		ORDER BY id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var tasks []storage.Task
	for rows.Next() {
		var t storage.Task
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

//Запрос задачи по ID.
func (s *Storage) GetById(id uint64) (*storage.Task, error) {
	var err error
	var t storage.Task
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
func (s *Storage) GetByAuthor(p interface{}) ([]storage.Task, error) {
	var rows pgx.Rows
	var err error
	head := `
	SELECT 
		tasks.id,
		tasks.opened,
		tasks.closed,
		tasks.author_id,
		tasks.assigned_id,
		tasks.title,
		tasks.content
	FROM tasks, users	
	`
	tail := `
	ORDER BY id;
	`
	switch p := p.(type) {
	case uint64:
		rows, err = s.db.Query(context.Background(),
			head+"WHERE tasks.author_id = $1 AND users.id = tasks.author_id"+tail,
			p,
		)
		if err != nil {
			return nil, err
		}
	case string:
		rows, err = s.db.Query(context.Background(),
			head+"WHERE users.name = $1 AND users.id = tasks.author_id"+tail,
			p,
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid type of query parameter")
	}
	var tasks []storage.Task
	for rows.Next() {
		var t storage.Task
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
func (s *Storage) GetByLabel(p interface{}) ([]storage.Task, error) {
	var rows pgx.Rows
	var err error
	head := `
		SELECT tasks.id, tasks.opened, tasks.closed, tasks.author_id,
			tasks.assigned_id, tasks.title, tasks.content
		FROM tasks, tasks_labels, labels
	`
	tail := `
		AND tasks_labels.task_id = tasks.id
		AND tasks_labels.label_id = labels.id
		ORDER BY tasks.id;
	`
	switch p := p.(type) {
	case uint64:
		rows, err = s.db.Query(context.Background(),
			head+"WHERE labels.id = $1"+tail,
			p,
		)
		if err != nil {
			return nil, err
		}
	case string:
		rows, err = s.db.Query(context.Background(),
			head+"WHERE labels.name = $1"+tail,
			p,
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid type of query parameter")
	}
	var tasks []storage.Task
	for rows.Next() {
		var t storage.Task
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
