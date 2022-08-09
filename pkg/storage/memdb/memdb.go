package memdb

import (
	"github.com/serjbibox/SF30.8.1/pkg/models"
)

type DB []models.Task

func NewMemDb() DB {
	return DB{}
}
