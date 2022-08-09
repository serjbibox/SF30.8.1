package postgresql

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DB_USERNAME = "serj"
	DB_HOST     = "192.168.52.129"
	DB_PORT     = "5432"
	DB_NAME     = "serjdb"
	DB_SSLMODE  = "require"
)

var elog = log.New(os.Stderr, "postgresql error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "postgresql info\t", log.Ldate|log.Ltime)

func GetConnectionString() string {
	pwd := os.Getenv("DbPass")
	if pwd == "" {
		elog.Fatalf("error reading password from os environment")
	}
	return "postgres://" +
		DB_USERNAME + ":" +
		pwd + "@" +
		DB_HOST + ":" +
		DB_PORT + "/" +
		DB_NAME + "?sslmode=" +
		DB_SSLMODE
}

func NewPostgresDB(constr string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, constr)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	ilog.Println("connected to postgres database")
	return db, nil
}
