package main

import (
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/SF30.8.1/pkg/storage"
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
	/*
		err = db.Delete(uint64(6))
		if err != nil {
			elog.Fatalf("error deleting database: %s", err.Error())
		}*/
	task, err := s.GetById(uint64(3))
	if err != nil {
		elog.Fatalf("error deleting database: %s", err.Error())
	}
	ilog.Println(task)

}
