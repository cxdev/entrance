package entrance

import (
	"entrance/backend/job"
	"entrance/backend/platform"
	"time"
)

type JobService struct {
	*platform.DB
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
	return nil
}
