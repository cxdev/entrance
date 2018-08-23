package entrance

import (
	"entrance/backend/exec"
	"entrance/backend/job"
	"entrance/backend/platform"
	"time"
)

type JobService struct {
	*platform.DB
	exec.ExecContextBuilder
}

type JobQuery struct {
	UpdatedFrom time.Time
	Updatedto   time.Time
	JobID       int
	JobStatus   int
	CommandQuery
}

func (s *JobService) Job(jobId int) *job.Job {
	return nil
}

func (s *JobService) Jobs(query *JobQuery) *[]job.Job {
	return nil
}

func (s *JobService) SaveJob(job *job.Job) (jobID int) {
	return -1
}

func (s *JobService) UpdateJob(job *job.Job, jobID int) error {
	return nil
}

func (s *JobService) RemoveJob(jobID int) error {
	return nil
}

func (s *JobService) ExecuteJob(job *job.Job) error {
	jobTag := string(job.Id)
	sysCmd := job.SysCmd

	execContext := s.CreateContext(jobTag, sysCmd)
	execCmd, err := execContext.CreateExecCmd()
	if err != nil {
		return err
	}

	err = execCmd.Start()
	if err != nil {
		return err
	}

	// TODO: consider for wait and kill case
	// cmd.Wait()
	// cmd.Process.Kill()
	return nil
}
