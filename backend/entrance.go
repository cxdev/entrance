package entrance

import (
	"entrance/backend/command"
	"entrance/backend/job"
)

type App struct {
	CommandService
	JobService
}

func (app *App) AddCommand(name string, commandType command.CommandType, segments string) int64 {
	command, _ := command.New(name, commandType, segments)
	result, err := app.SaveCommand(command)
	if err != nil {
		return -1
	}

	cid, err := result.LastInsertId()
	if err != nil {
		return -1
	}

	return cid
}

func (app *App) AddJob(cid int64, arguments string) int {
	command := app.Command(cid)
	job, _ := job.New(command, arguments)
	jobID := app.SaveJob(job)
	go app.ExecuteJob(job)
	return jobID
}
