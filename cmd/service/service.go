package main

import (
	"log"
	"os"

	"github.com/serjbibox/SF30.8.1/pkg/storage"
	"github.com/serjbibox/SF30.8.1/pkg/storage/postgresql"
)

var db storage.TaskStorage
var elog = log.New(os.Stderr, "service error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "service info\t", log.Ldate|log.Ltime)

func main() {
	var err error
	db, err = postgresql.New(postgresql.GetConnectionString())
	if err != nil {
		elog.Fatalf("error connecting database: %s", err.Error())
	}

	err = db.Delete(uint64(6))
	if err != nil {
		elog.Fatalf("error deleting database: %s", err.Error())
	}
	task, err := db.GetById(uint64(3))
	if err != nil {
		elog.Fatalf("error deleting database: %s", err.Error())
	}
	ilog.Println(task)

}
