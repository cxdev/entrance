package job

import (
	"entrance/backend/command"
	"entrance/backend/platform"
	"errors"
	"strings"
)

type Job struct {
	platform.BaseEntity
	Status    JobStatus
	CommandID int64
	Arguments *Arguments
	SysCmd    string
}

func New(command *command.Command, arguments string) (*Job, error) {
	arg, err := NewArguments(arguments)

	if err != nil {
		return nil, err
	}

	sysCmd, err := CreateSysCmd(command, arg)

	if err != nil {
		return nil, err
	}

	return &Job{platform.BaseEntity{}, WAITING, command.Id, arg, sysCmd}, nil
}

func CreateSysCmd(command *command.Command, arguments *Arguments) (string, error) {
	var sb strings.Builder
	for _, segment := range command.CommandSegments {
		if segment.IsRequired {
			sb.WriteString(segment.Base)
			sb.WriteString(" ")
			if segment.IsValuable {
				if argVal, ok := arguments.Get(segment.Key); ok {
					sb.WriteString(argVal)
					sb.WriteString(" ")
				} else {
					return "", errors.New("Not found error")
				}
			}
		} else {
			if argVal, ok := arguments.Get(segment.Key); ok {
				sb.WriteString(segment.Base)
				sb.WriteString(" ")
				if argVal != "" {
					sb.WriteString(argVal)
					sb.WriteString(" ")
				}
			}
		}
	}
	return strings.TrimSpace(sb.String()), nil
}
