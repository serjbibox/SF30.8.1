package postgresql

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var elog = log.New(os.Stderr, "postgresql error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "postgresql info\t", log.Ldate|log.Ltime)

func GetConnectionString() string {

	pwd := os.Getenv("DbPass")
	if pwd == "" {
		elog.Fatalf("error reading password from os environment")
	}
	return "postgres://" +
		viper.GetString("db.username") + ":" +
		pwd + "@" +
		viper.GetString("db.host") + ":" +
		viper.GetString("db.port") + "/" +
		viper.GetString("db.name") + "?sslmode=" +
		viper.GetString("db.sslmode")
}
