package storage

import (
	entrance "entrance/backend"
	"reflect"
	"testing"
	"time"
)

type EntityTestData struct {
	time1           time.Time
	time2           time.Time
	emptyBase       entrance.Base
	emptyModel      Model
	generalBase     entrance.Base
	generalModel    Model
	commands        []entrance.Command
	commandEntities []CommandEntity
	jobs            []entrance.Job
	jobEntities     []JobEntity
}

func NewEntityTestData() *EntityTestData {
	time1 := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	time2 := time.Now()
	emptyBase := entrance.Base{}
	emptyModel := Model{}
	generalBase := entrance.Base{1, time1, time2}
	generalModel := Model{1, time1, time2, nil}

	commandSegments1 := []entrance.CommandSegment{{"k1", "b1", false, true}}
	commandSegmentsString1 := `[{"Key":"k1","Base":"b1","IsRequired":false,"IsValuable":true}]`
	command1 := entrance.Command{generalBase, "test command 1", entrance.BATCH, commandSegments1}
	commandEntity1 := CommandEntity{generalModel, "test command 1", int(entrance.BATCH), commandSegmentsString1}

	commandSegments2 := []entrance.CommandSegment{{"k1", "b1", false, false}, {"k2", "b2", true, true}}
	commandSegmentsString2 := `[{"Key":"k1","Base":"b1","IsRequired":false,"IsValuable":false},{"Key":"k2","Base":"b2","IsRequired":true,"IsValuable":true}]`
	command2 := entrance.Command{generalBase, "test command 2", entrance.BATCH, commandSegments2}
	commandEntity2 := CommandEntity{generalModel, "test command 2", int(entrance.BATCH), commandSegmentsString2}

	commands := []entrance.Command{command1, command2}
	commandEntities := []CommandEntity{commandEntity1, commandEntity2}

	arguments1 := entrance.Arguments{"k1": "v1"}
	arguments2 := entrance.Arguments{"k1": "v1", "k2": "v2"}
	job1 := entrance.Job{generalBase, entrance.WAITING, 1, &arguments1, "ls /tmp"}
	jobEntity1 := JobEntity{generalModel, int(entrance.WAITING), 1, `{"k1":"v1"}`, "ls /tmp"}
	job2 := entrance.Job{generalBase, entrance.WAITING, 2, &arguments2, "ls -l /tmp"}
	jobEntity2 := JobEntity{generalModel, int(entrance.WAITING), 2, `{"k1":"v1","k2":"v2"}`, "ls -l /tmp"}

	jobs := []entrance.Job{job1, job2}
	jobEntities := []JobEntity{jobEntity1, jobEntity2}

	return &EntityTestData{time1, time2, emptyBase, emptyModel, generalBase, generalModel, commands, commandEntities, jobs, jobEntities}
}

func TestNewModel(t *testing.T) {
	testData := NewEntityTestData()
	type args struct {
		base *entrance.Base
	}
	tests := []struct {
		name    string
		args    args
		want    *Model
		wantErr bool
	}{
		{"nil case", args{nil}, nil, true},
		{"empty case", args{&testData.emptyBase}, &testData.emptyModel, false},
		{"general 1", args{&testData.commands[0].Base}, &testData.commandEntities[0].Model, false},
		{"general 2", args{&testData.commands[1].Base}, &testData.commandEntities[1].Model, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewModel(tt.args.base)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModel_ToBase(t *testing.T) {
	testData := NewEntityTestData()
	tests := []struct {
		name    string
		model   *Model
		want    *entrance.Base
		wantErr bool
	}{
		{"nil model", nil, nil, true},
		{"general mode", &Model{1, testData.time1, testData.time2, nil},
			&entrance.Base{1, testData.time1, testData.time2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.model.ToBase()
			if (err != nil) != tt.wantErr {
				t.Errorf("Model.ToBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Model.ToBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCommandEntity(t *testing.T) {
	testData := NewEntityTestData()
	type args struct {
		command *entrance.Command
	}
	tests := []struct {
		name    string
		args    args
		want    *CommandEntity
		wantErr bool
	}{
		{"nil command", args{nil}, nil, true},
		{"general command 1", args{&testData.commands[0]}, &testData.commandEntities[0], false},
		{"general command 2", args{&testData.commands[1]}, &testData.commandEntities[1], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommandEntity(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCommandEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommandEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandEntity_ToCommand(t *testing.T) {
	testData := NewEntityTestData()
	type fields struct {
		*CommandEntity
	}
	tests := []struct {
		name    string
		fields  fields
		want    *entrance.Command
		wantErr bool
	}{
		{"nil commandEntity", fields{nil}, nil, true},
		{"general commandEntity 1", fields{&testData.commandEntities[0]}, &testData.commands[0], false},
		{"general commandEntity 2", fields{&testData.commandEntities[1]}, &testData.commands[1], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.ToCommand()
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandEntity.ToCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandEntity.ToCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCommands(t *testing.T) {
	testData := NewEntityTestData()
	type args struct {
		commandEntities *[]CommandEntity
	}
	tests := []struct {
		name    string
		args    args
		want    *[]entrance.Command
		wantErr bool
	}{
		{"nil commands", args{nil}, nil, true},
		{"general commands", args{&testData.commandEntities}, &testData.commands, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCommands(tt.args.commandEntities)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJobEntity(t *testing.T) {
	testData := NewEntityTestData()
	type args struct {
		job *entrance.Job
	}
	tests := []struct {
		name    string
		args    args
		want    *JobEntity
		wantErr bool
	}{
		{"nil job", args{nil}, nil, true},
		{"general job 1", args{&testData.jobs[0]}, &testData.jobEntities[0], false},
		{"general job 2", args{&testData.jobs[1]}, &testData.jobEntities[1], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJobEntity(tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJobEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJobEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobEntity_ToJob(t *testing.T) {
	testData := NewEntityTestData()
	type fields struct {
		*JobEntity
	}
	tests := []struct {
		name    string
		fields  fields
		want    *entrance.Job
		wantErr bool
	}{
		{"nil jobEntity", fields{nil}, nil, true},
		{"general jobEntity 1", fields{&testData.jobEntities[0]}, &testData.jobs[0], false},
		{"general jobEntity 2", fields{&testData.jobEntities[1]}, &testData.jobs[1], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.ToJob()
			if (err != nil) != tt.wantErr {
				t.Errorf("JobEntity.ToJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JobEntity.ToJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToJobs(t *testing.T) {
	testData := NewEntityTestData()
	type args struct {
		jobEntities *[]JobEntity
	}
	tests := []struct {
		name    string
		args    args
		want    *[]entrance.Job
		wantErr bool
	}{
		{"nil jobs", args{nil}, nil, true},
		{"general commands", args{&testData.jobEntities}, &testData.jobs, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJobs(tt.args.jobEntities)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
