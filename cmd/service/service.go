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

	err = db.Delete(9)
	if err != nil {
		elog.Fatalf("error deleting database: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
