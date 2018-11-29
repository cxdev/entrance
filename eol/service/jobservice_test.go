package service

import (
	"entrance/eol/job"
	"entrance/eol/platform"
	"reflect"
	"testing"
)

func setupJobServiceTest() (*platform.DB, []*job.Job) {
	var TestDB, _ = platform.CreateDB("file::memory:?mode=memory")
	var TestJob1 = &job.Job{platform.BaseEntity{}, job.WAITING, 1, &job.Arguments{"k1": "v1"}, "testSystemCmd1"}
	var TestJob2 = &job.Job{platform.BaseEntity{}, job.RUNNING, 2, &job.Arguments{"k2": "v2"}, "testSystemCmd2"}
	var TestJob3 = &job.Job{platform.BaseEntity{}, job.DONE, 3, &job.Arguments{"k3": "v3"}, "testSystemCmd3"}
	return TestDB, []*job.Job{
		TestJob1, TestJob2, TestJob3,
	}
}

func TestJobService_Job(t *testing.T) {
	type fields struct {
		DB                 *platform.DB
		ExecContextBuilder exec.ExecContextBuilder
	}
	type args struct {
		jobId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *job.Job
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &JobService{
				DB:                 tt.fields.DB,
				ExecContextBuilder: tt.fields.ExecContextBuilder,
			}
			if got := s.Job(tt.args.jobId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JobService.Job() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobService_Jobs(t *testing.T) {
	type fields struct {
		DB                 *platform.DB
		ExecContextBuilder exec.ExecContextBuilder
	}
	type args struct {
		queryCondition *platform.QueryCondition
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]*job.Job
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &JobService{
				DB:                 tt.fields.DB,
				ExecContextBuilder: tt.fields.ExecContextBuilder,
			}
			if got := s.Jobs(tt.args.queryCondition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JobService.Jobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobService_SaveJob(t *testing.T) {
	type fields struct {
		DB                 *platform.DB
		ExecContextBuilder exec.ExecContextBuilder
	}
	type args struct {
		job *job.Job
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &JobService{
				DB:                 tt.fields.DB,
				ExecContextBuilder: tt.fields.ExecContextBuilder,
			}
			if got := s.SaveJob(tt.args.job); got != tt.want {
				t.Errorf("JobService.SaveJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobService_UpdateJob(t *testing.T) {
	type fields struct {
		DB                 *platform.DB
		ExecContextBuilder exec.ExecContextBuilder
	}
	type args struct {
		job   *job.Job
		jobID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &JobService{
				DB:                 tt.fields.DB,
				ExecContextBuilder: tt.fields.ExecContextBuilder,
			}
			if err := s.UpdateJob(tt.args.job, tt.args.jobID); (err != nil) != tt.wantErr {
				t.Errorf("JobService.UpdateJob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJobService_ExecuteJob(t *testing.T) {
	type fields struct {
		DB                 *platform.DB
		ExecContextBuilder exec.ExecContextBuilder
	}
	type args struct {
		jobID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &JobService{
				DB:                 tt.fields.DB,
				ExecContextBuilder: tt.fields.ExecContextBuilder,
			}
			if err := s.ExecuteJob(tt.args.jobID); (err != nil) != tt.wantErr {
				t.Errorf("JobService.ExecuteJob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
