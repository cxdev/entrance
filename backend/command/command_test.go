package command

import (
	"entrance/backend/platform"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name        string
		commandType CommandType
		segments    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Command
		wantErr bool
	}{
		{"test 1", args{"ls", BATCH, `[{"Key":"k1","Base":"b1","IsRequired":false,"IsValuable":false}]`},
			&Command{platform.BaseEntity{}, "ls", BATCH, CommandSegments{{"k1", "b1", false, false}}}, false},
		{"test 2", args{"ls", BATCH, `wrong`}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.name, tt.args.commandType, tt.args.segments)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
