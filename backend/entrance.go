package entrance

import (
	"entrance/backend/command"
	"entrance/backend/job"
)

type App struct {
	CommandService
	JobService
}

func (app *App) AddCommand(name string, commandType command.CommandType, segments string) int {
	command, _ := command.New(name, commandType, segments)
	return app.SaveCommand(command)
}

func (app *App) AddJob(cid int, arguments string) int {
	command := app.Command(cid)
	job, _ := job.New(command, arguments)
	jobID := app.SaveJob(job)
	go app.ExecuteJob(job)
	return jobID
}
