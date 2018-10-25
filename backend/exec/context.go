package exec

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type ExecContext struct {
	LogBase    string
	OutputPath string
	Errorpath  string
	SysCmd     string
}

type ExecContextBuilder struct {
	LogRoot string
}

type ResultData struct {
	OutputData *[]string
	ErrorData  *[]string
}

func (builder *ExecContextBuilder) CreateContext(jobTag string, sysCmd string) *ExecContext {
	return &ExecContext{
		path.Join(builder.LogRoot, jobTag),
		path.Join(builder.LogRoot, jobTag, "output.log"),
		path.Join(builder.LogRoot, jobTag, "error.log"),
		sysCmd,
	}
}

func (context *ExecContext) ExecCommand() error {
	cmdSlice := strings.Split(context.SysCmd, " ")
	cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)

	err := os.MkdirAll(context.LogBase, os.ModePerm)
	if err != nil {
		return err
	}

	outfile, err := os.Create(context.OutputPath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	errorfile, err := os.Create(context.Errorpath)
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

func (context *ExecContext) ReadResult() (*ResultData, error) {
	outputBytes, err := ioutil.ReadFile(context.OutputPath)
	if err != nil {
		return nil, err
	}

	errorBytes, err := ioutil.ReadFile(context.Errorpath)
	if err != nil {
		return nil, err
	}

	outputData := strings.Split(string(outputBytes), "\n")
	errorData := strings.Split(string(errorBytes), "\n")
	return &ResultData{&outputData, &errorData}, nil
}
