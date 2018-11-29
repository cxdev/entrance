package execute

import (
	entrance "entrance/backend"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type ExecuteService struct {
	LogRoot string
}

type ExecTask struct {
	LogBase    string
	OutputPath string
	Errorpath  string
	SysCmd     string
}

func NewExecuteService(logRoot string) *ExecuteService {
	return &ExecuteService{logRoot}
}

func (service *ExecuteService) Execute(job *entrance.Job) error {
	// TODO: if no ID or status is not waiting
	execTask := service.GetExecTask(getJobTag(job), job.SysCmd)
	return execTask.Execute()
}
func (service *ExecuteService) GetResult(job *entrance.Job) (*entrance.ExecuteResult, error) {
	execTask := service.GetExecTask(getJobTag(job), job.SysCmd)
	return execTask.ReadResult()
}

func (service *ExecuteService) GetExecTask(jobTag string, sysCmd string) *ExecTask {
	root := service.LogRoot
	return &ExecTask{
		path.Join(root, jobTag),
		path.Join(root, jobTag, "output.log"),
		path.Join(root, jobTag, "error.log"),
		sysCmd,
	}
}

func (execTask *ExecTask) Execute() error {
	cmdSlice := strings.Split(execTask.SysCmd, " ")
	cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)

	err := os.MkdirAll(execTask.LogBase, os.ModePerm)
	if err != nil {
		return err
	}

	outfile, err := os.Create(execTask.OutputPath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	errorfile, err := os.Create(execTask.Errorpath)
	if err != nil {
		return err
	}
	defer errorfile.Close()

	cmd.Stdout = outfile
	cmd.Stderr = errorfile

	err = cmd.Start()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (execTask *ExecTask) ReadResult() (*entrance.ExecuteResult, error) {
	outputBytes, err := ioutil.ReadFile(execTask.OutputPath)
	if err != nil {
		return nil, err
	}

	errorBytes, err := ioutil.ReadFile(execTask.Errorpath)
	if err != nil {
		return nil, err
	}

	outputData := strings.Split(string(outputBytes), "\n")
	errorData := strings.Split(string(errorBytes), "\n")
	return &entrance.ExecuteResult{&outputData, &errorData}, nil
}

func getJobTag(job *entrance.Job) string {
	return strconv.FormatUint(uint64(job.ID), 10)
}
