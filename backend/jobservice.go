package entrance

import (
	"entrance/backend/exec"
	"entrance/backend/job"
	"entrance/backend/platform"
	"errors"
	"log"
	"strconv"
	"time"
)

type JobService struct {
	*platform.DB
	exec.ExecContextBuilder
}

func (s *JobService) Job(jobId int64) *job.Job {
	qc := platform.QueryCondition{"id": jobId}

	jobs := s.Jobs(&qc)
	if jobs != nil && len(*jobs) > 0 {
		return (*jobs)[0]
	}
	return nil
}

func (s *JobService) Jobs(queryCondition *platform.QueryCondition) *[]*job.Job {
	var (
		id        int64
		status    job.JobStatus
		commandID int64
		arguments string
		systemCmd string
		created   time.Time
		updated   time.Time
	)

	sqlBase := "SELECT * FROM jobs"

	rows, err := s.QueryByCondition(sqlBase, queryCondition)
	if err != nil {
		return nil
	}
	var jobs []*job.Job
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &status, &commandID, &arguments, &systemCmd, &created, &updated)
		// TODO: consider about usage of &commandID and &arguments
		job := &job.Job{platform.BaseEntity{}, status, commandID, nil, systemCmd}
		job.SetupEntity(id, created, updated)
		jobs = append(jobs, job)
	}

	return &jobs
}

func (s *JobService) SaveJob(job *job.Job) int64 {
	sql := "INSERT INTO jobs(status, command_id, arguments, system_cmd) VALUES(?, ?, ?, ?)"
	arguments, err := job.Arguments.ToString()
	if err != nil {
		return -1
	}
	result, err := s.Exec(sql, job.Status, job.CommandID, arguments, job.SysCmd)
	if err != nil {
		return -1
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1
	}
	return id
}

func (s *JobService) UpdateJob(job *job.Job, jobID int) error {
	sql := "UPDATE jobs SET status=?, command_id=?, arguments=?, system_cmd=? WHERE id=?"
	arguments, err := job.Arguments.ToString()
	if err != nil {
		return err
	}
	_, err = s.Exec(sql, job.Status, job.CommandID, arguments, job.SysCmd, job.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *JobService) RemoveJob(jobID int) error {
	return nil
}

func (s *JobService) ExecuteJob(jobID int64) error {
	job := s.Job(jobID)
	jobTag := strconv.FormatInt(jobID, 10)
	sysCmd := job.SysCmd

	execContext := s.CreateContext(jobTag, sysCmd)
	err := execContext.ExecCommand()
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Done run command")

	// TODO: consider for wait and kill case
	// cmd.Wait()
	// cmd.Process.Kill()
	return nil
}

func (s *JobService) ReadJobResult(jobID int64) (*exec.ResultData, error) {
	job := s.Job(jobID)
	if job == nil {
		return nil, errors.New("Not found the request job")
	}
	jobTag := strconv.FormatInt(jobID, 10)
	execContext := s.CreateContext(jobTag, job.SysCmd)
	return execContext.ReadResult()
}
