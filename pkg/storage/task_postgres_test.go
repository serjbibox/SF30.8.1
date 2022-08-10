package storage

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/serjbibox/SF30.8.1/pkg/models"
	"github.com/serjbibox/SF30.8.1/pkg/storage/postgresql"
)

var s *Storage

func TestMain(m *testing.M) {
	db, err := postgresql.NewPostgresDB(postgresql.GetConnectionString())
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}
	s = NewStoragePostgres(db)
	os.Exit(m.Run())
}

func TestTaskPostgres_GetAll(t *testing.T) {
	data, err := s.GetAll()
	if err != nil {
		t.Errorf("TestStorage_GetAll error = %v", err)
	}
	t.Log(data)
}

func TestTaskPostgres_Create(t *testing.T) {
	task := models.Task{
		Content: "test task",
	}
	id, err := s.Create(task)
	if err != nil {
		t.Errorf("TestStorage_Create error = %v", err)
	}
	t.Log(id)
}

func TestTaskPostgres_Update(t *testing.T) {
	type args struct {
		id uint64
		t  models.Task
	}
	ar := args{
		id: 3,
		t: models.Task{
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

func TestTaskPostgres_Delete(t *testing.T) {
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

func TestTaskPostgres_buildLabelQuery(t *testing.T) {
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
		{
			name: "buildLabelQuery_3 wrong type of argument",
			args: args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildLabelQuery(tt.args.t)
			t.Log(got)
		})
	}
}

func TestTaskPostgres_buildAuthorQuery(t *testing.T) {
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
		{
			name: "buildLabelQuery_3 wrong type of argument",
			args: args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildAuthorQuery(tt.args.t)
			t.Log(got)
		})
	}
}

func TestTaskPostgres_GetById(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetById",
			s:    s,
			args: args{2},
			want: &models.Task{2, 0, 0, 0, 0, "updated task", "task updated"},
		},
		{
			name:    "GetById with error",
			s:       s,
			args:    args{0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskPostgres.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskPostgres.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskPostgres_GetByAuthor(t *testing.T) {
	type args struct {
		p interface{}
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    []models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetByAuthor id",
			s:    s,
			args: args{uint64(2)},
			want: []models.Task{
				{8, 1659697344, 0, 2, 0, "task", "do it"},
			},
		},
		{
			name: "GetByAuthor name",
			s:    s,
			args: args{"Petya"},
			want: []models.Task{
				{8, 1659697344, 0, 2, 0, "task", "do it"},
			},
		},
		{
			name:    "GetByAuthor vis error invalid type param",
			s:       s,
			args:    args{100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetByAuthor(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskPostgres.GetByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskPostgres.GetByAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskPostgres_GetByLabel(t *testing.T) {
	type args struct {
		p interface{}
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    []models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetByLabel id",
			s:    s,
			args: args{uint64(3)},
			want: []models.Task{
				{5, 1659694979, 0, 0, 0, "go", "gogo"},
			},
		},
		{
			name: "GetByLabel name",
			s:    s,
			args: args{"task3"},
			want: []models.Task{
				{5, 1659694979, 0, 0, 0, "go", "gogo"},
			},
		},
		{
			name:    "GetByLabel vis error invalid type param",
			s:       s,
			args:    args{100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetByLabel(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskPostgres.GetByLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskPostgres.GetByLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}
