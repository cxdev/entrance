package job

import (
	"entrance/backend/command"
	"entrance/backend/platform"
)

type Job struct {
	platform.BaseEntity
	Status    JobStatus
	Command   *command.Command
	Arguments *Arguments
}

func New(command *command.Command, arguments string) (*Job, error) {
	arg, err := NewArguments(arguments)

	if err != nil {
		return nil, err
	}

	return &Job{platform.BaseEntity{}, WAITING, command, arg}, nil
}
