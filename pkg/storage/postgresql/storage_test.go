package postgresql

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/serjbibox/SF30.8.1/pkg/storage"
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

func TestStorage_Create(t *testing.T) {
	task := storage.Task{
		Content: "test task",
	}
	id, err := s.Create(task)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestStorage_Update(t *testing.T) {
	type args struct {
		id uint64
		t  storage.Task
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Update(tt.args.id, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
