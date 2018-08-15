package command

import (
	"reflect"
	"testing"
)

type testData struct {
	segments CommandSegments
	json     string
}

func setup() []testData {
	return []testData{
		{CommandSegments{{"k1", "b1", false, false}}, `[{"Key":"k1","Base":"b1","IsRequired":false,"IsValuable":false}]`},
		{CommandSegments{{"k1", "b1", false, false}, {"k2", "b2", true, true}}, `[{"Key":"k1","Base":"b1","IsRequired":false,"IsValuable":false},{"Key":"k2","Base":"b2","IsRequired":true,"IsValuable":true}]`},
	}
}

func TestNewCommandSegments(t *testing.T) {
	testDatas := setup()
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *CommandSegments
		wantErr bool
	}{
		{"test 1", args{testDatas[0].json}, &testDatas[0].segments, false},
		{"test 2", args{testDatas[1].json}, &testDatas[1].segments, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommandSegments(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCommandSegments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommandSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandSegments_ToString(t *testing.T) {
	testDatas := setup()
	tests := []struct {
		name     string
		segments *CommandSegments
		want     string
		wantErr  bool
	}{
		{"test 1", &testDatas[0].segments, testDatas[0].json, false},
		{"test 2", &testDatas[1].segments, testDatas[1].json, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.segments.ToString()
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandSegments.ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CommandSegments.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandSegments_Load(t *testing.T) {
	testDatas := setup()
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *CommandSegments
		wantErr bool
	}{
		{"test 1", args{testDatas[0].json}, &testDatas[0].segments, false},
		{"test 2", args{testDatas[1].json}, &testDatas[1].segments, false},
	}
	for _, tt := range tests {
		segments := make(CommandSegments, 0)
		t.Run(tt.name, func(t *testing.T) {
			if err := segments.Load(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("CommandSegments.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(&segments, tt.want) {
				t.Errorf("CommandSegments.Load() = %v, want %v", segments, tt.want)
			}
		})
	}
}
