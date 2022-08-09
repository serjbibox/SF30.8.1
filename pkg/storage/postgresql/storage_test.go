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
	ar := args{
		id: 3,
		t: storage.Task{
			Content: "tested task",
		},
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Update_1",
			s:    s,
			want: ar.id,
			args: ar,
		},
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
			t.Log(got)
		})
	}
}

func TestStorage_Delete(t *testing.T) {
	type args struct {
		taskid uint64
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		{
			name: "delete_1",
			s:    s,
			args: args{
				taskid: uint64(7),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Delete(tt.args.taskid); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_buildLabelQuery(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "buildLabelQuery_1",
			args: args{1},
		},
		{
			name: "buildLabelQuery_2",
			args: args{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildLabelQuery(tt.args.t)
			t.Log(got)
		})
	}
}
