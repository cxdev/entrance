package entrance

import (
	"entrance/backend/command"
	"entrance/backend/platform"
	"reflect"
	"testing"
)

func setup() (*platform.DB, []*command.Command) {
	// var TestDB, _ = platform.CreateDB("/tmp/ut.db")
	var TestDB, _ = platform.CreateDB("file::memory:?mode=memory")
	var TestCommand1, _ = command.New("ls", command.BATCH, `[{"Key":"k1","Base":"b1","IsRequired":false,"IsValuable":false}]`)
	var TestCommand2, _ = command.New("cmd2", command.BATCH, `[{"Key":"k1","Base":"b1","IsRequired":false,"IsValuable":false}]`)
	var TestCommand3, _ = command.New("ls", command.BATCH, `[{"Key":"k1","Base":"b1","IsRequired":true,"IsValuable":false}]`)
	TestCommand3.Id = 1
	return TestDB, []*command.Command{
		TestCommand1, TestCommand2, TestCommand3,
	}
}

func TestCommandService_SaveCommand(t *testing.T) {
	var TestDB, testCommands = setup()
	type fields struct {
		DB *platform.DB
	}
	type args struct {
		command *command.Command
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantInsertID int64
		wantErr      bool
	}{
		{"test 1", fields{TestDB}, args{testCommands[0]}, 1, false},
		{"test 2", fields{TestDB}, args{testCommands[1]}, 2, false},
		{"test 3", fields{TestDB}, args{testCommands[2]}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommandService{
				DB: tt.fields.DB,
			}
			got, err := s.SaveCommand(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandService.SaveCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			insertID, _ := got.LastInsertId()
			if !reflect.DeepEqual(insertID, tt.wantInsertID) {
				t.Errorf("CommandService.SaveCommand()'s LastInsertId = %v, wantInsertID %v", insertID, tt.wantInsertID)
			}
		})
	}
}

func TestCommandService_Command(t *testing.T) {
	var TestDB, testCommands = setup()
	(&CommandService{TestDB}).SaveCommand(testCommands[0])
	(&CommandService{TestDB}).SaveCommand(testCommands[1])

	type fields struct {
		DB *platform.DB
	}
	type args struct {
		cid int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *command.Command
	}{
		{"test 1", fields{TestDB}, args{1}, testCommands[0]},
		{"test 2", fields{TestDB}, args{2}, testCommands[1]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommandService{
				DB: tt.fields.DB,
			}
			got := s.Command(tt.args.cid)
			// Ignore base entity
			got.BaseEntity = tt.want.BaseEntity
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}
