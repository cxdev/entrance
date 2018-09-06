package entrance

import (
	"database/sql"
	"entrance/backend/command"
	"entrance/backend/platform"
	"time"
)

type CommandService struct {
	*platform.DB
}

func (s *CommandService) Command(cid int64) *command.Command {
	qc := platform.QueryCondition{"id": cid}

	commands := s.Commands(&qc)
	if len(*commands) > 0 {
		return (*commands)[0]
	}
	return nil
}

func (s *CommandService) Commands(queryCondition *platform.QueryCondition) *[]*command.Command {
	var (
		id          int64
		name        string
		commandType command.CommandType
		segments    string
		created     time.Time
		updated     time.Time
	)

	sqlBase := "SELECT * FROM commands"

	rows, err := s.QueryByCondition(sqlBase, queryCondition)
	if err != nil {
		return nil
	}

	var commands []*command.Command
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &commandType, &segments, &created, &updated)
		command, _ := command.New(name, commandType, segments)
		command.SetupEntity(id, created, updated)
		commands = append(commands, command)
	}

	return &commands
}

func (s *CommandService) SaveCommand(command *command.Command) (sql.Result, error) {
	commandSegments, err := command.CommandSegments.ToString()
	if err != nil {
		return nil, err
	}
	var args = []interface{}{command.Name, command.CommandType, commandSegments}
	var sql string

	if command.IsEntity() {
		sql = "UPDATE commands SET name=?, command_type=?, command_segments=? WHERE id=?"
		args = append(args, command.Id)
	} else {
		sql = "INSERT INTO commands(name, command_type, command_segments) VALUES(?, ?, ?)"
	}

	return s.Exec(sql, args...)
}
