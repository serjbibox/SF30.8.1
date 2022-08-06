package main

import (
	"log"
	"os"

	"github.com/serjbibox/SF30.8.1/pkg/storage"
	"github.com/serjbibox/SF30.8.1/pkg/storage/postgresql"
	"github.com/spf13/viper"
)

var db storage.Storage
var elog = log.New(os.Stderr, "service error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "service info\t", log.Ldate|log.Ltime)

func main() {
	var err error
	err = initConfig()
	if err != nil {
		elog.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err = postgresql.New(postgresql.GetConnectionString())
	if err != nil {
		elog.Fatalf("error connecting database: %s", err.Error())
	}
	/*id, err := db.Create(storage.Task{
		AuthorID: 1,
		Title:    "task",
		Content:  "do it",
	})
	if err != nil {
		elog.Println(err)
	} else {
		ilog.Println("id: ", id)
	}*/
	//var authorId string = "Petya"
	/*var authorId uint32 = 1
	tasks, err := db.GetByAuthor(authorId)
	if err != nil {
		elog.Println(err)
	}*/

	var labelId string = "task2"
	tasks, err := db.GetByLabel(labelId)
	if err != nil {
		elog.Println(err)
	}
	for _, val := range tasks {
		ilog.Println("id:", val)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
