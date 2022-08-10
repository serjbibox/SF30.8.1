package main

import (
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/SF30.8.1/pkg/models"
	"github.com/serjbibox/SF30.8.1/pkg/storage"
	"github.com/serjbibox/SF30.8.1/pkg/storage/memdb"
	"github.com/serjbibox/SF30.8.1/pkg/storage/postgresql"
)

var db *pgxpool.Pool
var elog = log.New(os.Stderr, "service error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "service info\t", log.Ldate|log.Ltime)

func main() {
	var err error
	db, err = postgresql.NewPostgresDB(postgresql.GetConnectionString())
	if err != nil {
		elog.Fatalf("error connecting database: %s", err.Error())
	}
	s := storage.NewStoragePostgres(db)
	t, err := s.GetById(uint64(2))
	if err != nil {
		elog.Println(err)
	}
	ilog.Println(t)

	memdb := memdb.NewMemDb()
	s = storage.NewStorageMemDb(memdb)

	id, _ := s.Create(models.Task{
		ID:         1,
		Title:      "memdb task 1",
		Content:    "new task 1",
		AuthorID:   1,
		AssignedID: 1,
	})
	ilog.Println(id)
	id, _ = s.Create(models.Task{
		ID:         2,
		Title:      "memdb task 2",
		Content:    "new task 2",
		AuthorID:   2,
		AssignedID: 2,
	})
	ilog.Println(id)
	task, err := s.GetById(uint64(1))
	if err != nil {
		elog.Fatalf("%s", err.Error())
	}
	ilog.Println(task)

}
