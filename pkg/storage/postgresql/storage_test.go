package postgresql

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var s *Storage

func TestMain(m *testing.M) {
	var err error
	pwd := os.Getenv("DbPass")
	fmt.Println(pwd)
	if pwd == "" {
		m.Run()
	}
	s, err = New("postgres://serj:" +
		pwd + "@192.168.52.129:5432/serjdb?sslmode=require")
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}
	os.Exit(m.Run())
}

func TestStorage_GetAll(t *testing.T) {
	data, err := s.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
