package storage

import (
	"entrance/backend"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestNewStorageService(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    *StorageService
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStorageService(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStorageService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStorageService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_CreateCommand(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		name            string
		commandtype     entrance.CommandType
		commandSegments []entrance.CommandSegment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entrance.Command
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			got, err := service.CreateCommand(tt.args.name, tt.args.commandtype, tt.args.commandSegments)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageService.CreateCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageService.CreateCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_Command(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		cID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entrance.Command
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			got, err := service.Command(tt.args.cID)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageService.Command() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageService.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_Commands(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		qc *entrance.QueryCondition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]*entrance.Command
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			got, err := service.Commands(tt.args.qc)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageService.Commands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageService.Commands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_SaveCommand(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		command *entrance.Command
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			if err := service.SaveCommand(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("StorageService.SaveCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageService_CreateJob(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		cID       uint
		arguments *entrance.Arguments
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entrance.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			got, err := service.CreateJob(tt.args.cID, tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageService.CreateJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageService.CreateJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_Job(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		jID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entrance.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			got, err := service.Job(tt.args.jID)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageService.Job() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageService.Job() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_Jobs(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		qc *entrance.QueryCondition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]*entrance.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			got, err := service.Jobs(tt.args.qc)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageService.Jobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageService.Jobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_SaveJob(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		job *entrance.Job
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &StorageService{
				db: tt.fields.db,
			}
			if err := service.SaveJob(tt.args.job); (err != nil) != tt.wantErr {
				t.Errorf("StorageService.SaveJob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createSysCmd(t *testing.T) {
	type args struct {
		cs        *[]entrance.CommandSegment
		arguments *entrance.Arguments
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createSysCmd(tt.args.cs, tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSysCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createSysCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_separateConditionsWithValues(t *testing.T) {
	type args struct {
		qc *entrance.QueryCondition
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := separateConditionsWithValues(tt.args.qc)
			if got != tt.want {
				t.Errorf("separateConditionsWithValues() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("separateConditionsWithValues() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
