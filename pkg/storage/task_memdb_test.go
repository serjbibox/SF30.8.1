package storage

import (
	"reflect"
	"testing"

	"github.com/serjbibox/SF30.8.1/pkg/models"
	"github.com/serjbibox/SF30.8.1/pkg/storage/memdb"
)

func TestNewTaskMemDb(t *testing.T) {
	type args struct {
		db memdb.DB
	}
	tests := []struct {
		name string
		args args
		want TaskMemdb
	}{
		// TODO: Add test cases.
		{
			name: "new memdb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskMemDb(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskMemDb() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestTaskMemdb_Create(t *testing.T) {
	type args struct {
		t models.Task
	}
	tests := []struct {
		name    string
		s       *TaskMemdb
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "memdb create",
			s: &TaskMemdb{
				db: memdb.NewMemDb(),
			},
			args: args{
				models.Task{
					ID:      1,
					Title:   "task1",
					Content: "task1 content",
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Create(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskMemdb.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TaskMemdb.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskMemdb_Delete(t *testing.T) {
	type args struct {
		taskid uint64
	}
	s := &TaskMemdb{
		db: memdb.DB{
			models.Task{
				ID:      1,
				Title:   "task1",
				Content: "task1 content",
			},
		},
	}
	tests := []struct {
		name    string
		s       *TaskMemdb
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "memdb delete",
			s:       s,
			args:    args{1},
			wantErr: false,
		},
		{
			name:    "memdb delete with error id = 0",
			s:       s,
			args:    args{0},
			wantErr: true,
		},
		{
			name:    "memdb delete with error no matched entry",
			s:       s,
			args:    args{5},
			wantErr: true,
		},
		{
			name:    "memdb delete with error no entries",
			s:       &TaskMemdb{},
			args:    args{1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Delete(tt.args.taskid); (err != nil) != tt.wantErr {
				t.Errorf("TaskMemdb.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskMemdb_GetAll(t *testing.T) {
	s := &TaskMemdb{
		db: memdb.DB{
			models.Task{
				ID:      1,
				Title:   "task1",
				Content: "task1 content",
			},
		},
	}
	tests := []struct {
		name    string
		s       *TaskMemdb
		want    []models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "memdb GetAll",
			s:       s,
			want:    s.db,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskMemdb.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskMemdb.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskMemdb_Update(t *testing.T) {
	s := &TaskMemdb{
		db: memdb.DB{
			models.Task{
				ID:      1,
				Title:   "task1",
				Content: "task1 content",
			},
		},
	}
	type args struct {
		id uint64
		t  models.Task
	}
	tests := []struct {
		name    string
		s       *TaskMemdb
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "memdb Update",
			s:    s,
			args: args{
				id: uint64(1),
				t: models.Task{
					ID:      1,
					Title:   "task_updated",
					Content: "task1 content",
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Update(tt.args.id, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskMemdb.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TaskMemdb.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskMemdb_GetById(t *testing.T) {
	s := &TaskMemdb{
		db: memdb.DB{
			models.Task{
				ID:      1,
				Title:   "task1",
				Content: "task1 content",
			},
		},
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		s       *TaskMemdb
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "memdb GetById",
			s:    s,
			args: args{
				id: uint64(1),
			},
			want:    &s.db[0],
			wantErr: false,
		},
		{
			name: "memdb GetById with error id=0",
			s:    s,
			args: args{
				id: uint64(0),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "memdb GetById with error no matched entry",
			s:    s,
			args: args{
				id: uint64(5),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "memdb GetById with error no entries",
			s:    &TaskMemdb{},
			args: args{
				id: uint64(5),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskMemdb.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskMemdb.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskMemdb_GetByAuthor(t *testing.T) {
	s := &TaskMemdb{
		db: memdb.DB{
			models.Task{
				ID:       1,
				AuthorID: 1,
				Title:    "task1",
				Content:  "task1 content",
			},
		},
	}
	type args struct {
		p interface{}
	}
	tests := []struct {
		name    string
		s       *TaskMemdb
		args    args
		want    []models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "memdb GetByAuthor",
			s:    s,
			args: args{
				p: uint64(1),
			},
			want:    s.db,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetByAuthor(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskMemdb.GetByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskMemdb.GetByAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskMemdb_GetByLabel(t *testing.T) {
	s := &TaskMemdb{
		db: memdb.DB{
			models.Task{
				ID:       1,
				AuthorID: 1,
				Title:    "task1",
				Content:  "task1 content",
			},
		},
	}
	type args struct {
		p interface{}
	}
	tests := []struct {
		name    string
		s       *TaskMemdb
		args    args
		want    []models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "memdb GetByAuthor",
			s:    s,
			args: args{
				p: uint64(1),
			},
			want:    s.db,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetByLabel(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskMemdb.GetByLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskMemdb.GetByLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}
