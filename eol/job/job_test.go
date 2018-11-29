package job

import (
	"entrance/backend/command"
	"testing"
)

func TestCreateSysCmd(t *testing.T) {
	segments := command.CommandSegments{{Key: "ls", Base: "ls", IsRequired: true, IsValuable: true}}
	testSegments, _ := segments.ToString()
	testCommand, _ := command.New("ls", command.BATCH, testSegments)

	type args struct {
		command   *command.Command
		arguments *Arguments
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test 1", args{testCommand, &Arguments{"no_useful_key": "no_useful_val"}}, "", true},
		{"test 2", args{testCommand, &Arguments{"ls": ""}}, "ls", false},
		{"test 2", args{testCommand, &Arguments{"ls": "/tmp"}}, "ls /tmp", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSysCmd(tt.args.command, tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSysCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSysCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
