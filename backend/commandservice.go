package entrance

import (
	"entrance/backend/command"
	"entrance/backend/platform"
	"time"
)

type CommandService struct {
	*platform.DB
}

type CommandQuery struct {
	UpdatedFrom time.Time
	Updatedto   time.Time
	CommandId   int
	ProcessType int
}

func (s *CommandService) Command(cid int) *command.Command {
	return nil
}

func (s *CommandService) Commands(query *CommandQuery) *[]command.Command {
	return nil
}

func (s *CommandService) SaveCommand(command *command.Command) (cid int) {
	return -1
}
