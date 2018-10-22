package entrance

import (
	"entrance/backend/command"
	"entrance/backend/job"
)

type App struct {
	CommandService
	JobService
}

func (app *App) AddCommand(name string, commandType command.CommandType, segments string) (int64, error) {
	command, err := command.New(name, commandType, segments)
	if err != nil {
		return -1, err
	}

	result, err := app.SaveCommand(command)
	if err != nil {
		return -1, err
	}

	cid, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return cid, nil
}

func (app *App) AddJob(cid int64, arguments string) (int64, error) {
	command := app.Command(cid)
	if command == nil {
		return -1, nil
	}

	job, err := job.New(command, arguments)
	if err != nil {
		return -1, err
	}

	jobID := app.SaveJob(job)
	go app.ExecuteJob(jobID)
	return jobID, nil
}
