package storage

// import (
// 	"entrance/backend"
// 	"reflect"
// 	"testing"

// 	"github.com/jinzhu/gorm"
// )

// type CommandSegment entrance.CommandSegment

// func TestEntranceService_CreateCommand(t *testing.T) {

// 	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
// 	defer db.Close()

// 	commandSegments := []entrance.CommandSegment{entrance.CommandSegment{"k1", "b1", false, false}, entrance.CommandSegment{"k2", "b2", true, false}}

// 	type fields struct {
// 		DB *gorm.DB
// 	}
// 	type args struct {
// 		name            string
// 		commandtype     entrance.CommandType
// 		commandSegments []entrance.CommandSegment
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *entrance.Command
// 		wantErr bool
// 	}{
// 		{"test 1", fields{db}, args{"test", entrance.BATCH, commandSegments}, nil, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			service, _ := NewEntranceService(db)
// 			got, err := service.CreateCommand(tt.args.name, tt.args.commandtype, tt.args.commandSegments)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("EntranceService.CreateCommand() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("EntranceService.CreateCommand() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
