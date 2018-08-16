package job

import (
	"reflect"
	"testing"
)

func testArguments() []Arguments {
	return []Arguments{
		map[string]string{"k1": "v1"},
		map[string]string{"k1": "v1", "k2": "v2"},
	}
}

func TestNewArguments(t *testing.T) {
	argsCases := testArguments()

	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *Arguments
		wantErr bool
	}{
		{"test 1", args{"{\"k1\":\"v1\"}"}, &argsCases[0], false},
		{"test 2", args{"{\"k1\":\"v1\",\"k2\":\"v2\"}"}, &argsCases[1], false},
		{"error case", args{"wrong json"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewArguments(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewArguments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArguments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArguments_ToString(t *testing.T) {
	argsCases := testArguments()
	tests := []struct {
		name      string
		arguments *Arguments
		want      string
		wantErr   bool
	}{
		{"test 1", &argsCases[0], "{\"k1\":\"v1\"}", false},
		{"test 2", &argsCases[1], "{\"k1\":\"v1\",\"k2\":\"v2\"}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.arguments.ToString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Arguments.ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Arguments.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArguments_Load(t *testing.T) {
	argsCases := testArguments()
	argument := make(Arguments)

	type args struct {
		arguments *Arguments
		str       string
	}
	tests := []struct {
		name    string
		args    args
		want    *Arguments
		wantErr bool
	}{
		{"test 1", args{&argument, "{\"k1\":\"v1\"}"}, &argsCases[0], false},
		{"test 2", args{&argument, "{\"k1\":\"v1\",\"k2\":\"v2\"}"}, &argsCases[1], false},
		{"test 3", args{&argument, "wrong json"}, &argument, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.arguments.Load(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Arguments.Load() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.arguments
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArguments_Get(t *testing.T) {
	argsCases := testArguments()

	type args struct {
		key string
	}
	tests := []struct {
		name      string
		arguments *Arguments
		args      args
		want      string
		want1     bool
	}{
		{"test 1", &argsCases[0], args{"k1"}, "v1", true},
		{"test 2", &argsCases[0], args{"k2"}, "", false},
		{"test 3", &argsCases[1], args{"k1"}, "v1", true},
		{"test 4", &argsCases[1], args{"k2"}, "v2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.arguments.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("Arguments.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Arguments.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
