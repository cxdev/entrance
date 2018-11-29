package storage

import (
	"encoding/json"
	entrance "entrance/backend"
	"errors"

	"github.com/jinzhu/gorm"
)

type Model gorm.Model
type CommandEntity struct {
	Model
	Name            string
	CommandType     int
	CommandSegments string
}

type JobEntity struct {
	Model
	Status    int
	CommandID uint
	Arguments string
	SysCmd    string
}

func NewModel(base *entrance.Base) (*Model, error) {
	if base == nil {
		return nil, errors.New("nil entity")
	}

	if (*base == entrance.Base{}) {
		return &Model{}, nil
	}

	return &Model{base.ID, base.CreatedAt, base.UpdatedAt, nil}, nil

}

func (model *Model) ToBase() (*entrance.Base, error) {
	if model == nil {
		return nil, errors.New("nil entity")
	}

	return &entrance.Base{model.ID, model.CreatedAt, model.UpdatedAt}, nil
}

func NewCommandEntity(command *entrance.Command) (*CommandEntity, error) {
	if command == nil {
		return nil, errors.New("nil entity")
	}

	bytes, err := json.Marshal(command.CommandSegments)
	if err != nil {
		return nil, err
	}

	model, err := NewModel(&command.Base)
	if err != nil {
		return nil, err
	}

	commandEntity := CommandEntity{*model, command.Name, int(command.CommandType), string(bytes)}
	return &commandEntity, nil
}

func (commandEntity *CommandEntity) ToCommand() (*entrance.Command, error) {
	if commandEntity == nil {
		return nil, errors.New("nil entity")
	}
	var commandSegments []entrance.CommandSegment
	err := json.Unmarshal([]byte(commandEntity.CommandSegments), &commandSegments)
	if err != nil {
		return nil, err
	}

	base, err := commandEntity.ToBase()
	if err != nil {
		return nil, err
	}

	command := entrance.Command{*base, commandEntity.Name, entrance.CommandType(commandEntity.CommandType), commandSegments}
	return &command, nil
}

func ToCommands(commandEntities *[]CommandEntity) (*[]entrance.Command, error) {
	if commandEntities == nil {
		return nil, errors.New("nil entity")
	}
	commands := make([]entrance.Command, len(*commandEntities))
	for i := range commands {
		command, err := (*commandEntities)[i].ToCommand()
		if err != nil {
			return nil, err
		}
		commands[i] = *command
	}
	return &commands, nil
}

func NewJobEntity(job *entrance.Job) (*JobEntity, error) {
	if job == nil {
		return nil, errors.New("nil entity")
	}

	bytes, err := json.Marshal(job.Arguments)
	if err != nil {
		return nil, err
	}
	model, err := NewModel(&job.Base)
	if err != nil {
		return nil, err
	}
	jobEntity := JobEntity{*model, int(job.Status), job.CommandID, string(bytes), job.SysCmd}
	return &jobEntity, nil
}

func (jobEntity *JobEntity) ToJob() (*entrance.Job, error) {
	if jobEntity == nil {
		return nil, errors.New("nil entity")
	}
	var arguments entrance.Arguments
	err := json.Unmarshal([]byte(jobEntity.Arguments), &arguments)
	if err != nil {
		return nil, err
	}

	base, err := jobEntity.ToBase()
	if err != nil {
		return nil, err
	}
	job := entrance.Job{*base, entrance.JobStatus(jobEntity.Status), jobEntity.CommandID, &arguments, jobEntity.SysCmd}
	return &job, nil
}

func ToJobs(jobEntities *[]JobEntity) (*[]entrance.Job, error) {
	if jobEntities == nil {
		return nil, errors.New("nil entity")
	}
	jobs := make([]entrance.Job, len(*jobEntities))
	for i := range jobs {
		job, err := (*jobEntities)[i].ToJob()
		if err != nil {
			return nil, err
		}
		jobs[i] = *job
	}
	return &jobs, nil
}
