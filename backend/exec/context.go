package exec

import (
	"os"
	"os/exec"
	"path"
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

func (builder *ExecContextBuilder) CreateContext(jobTag string, sysCmd string) *ExecContext {
	return &ExecContext{
		path.Join(builder.LogRoot, jobTag),
		path.Join(builder.LogRoot, jobTag, "output.log"),
		path.Join(builder.LogRoot, jobTag, "error.log"),
		sysCmd,
	}
}

func (context *ExecContext) CreateExecCmd() (*exec.Cmd, error) {
	cmd := exec.Command(context.SysCmd)

	outfile, err := os.Create(context.OutputPath)
	if err != nil {
		return nil, err
	}
	defer outfile.Close()

	errorfile, err := os.Create(context.Errorpath)
	if err != nil {
		return nil, err
	}
	defer errorfile.Close()

	cmd.Stdout = outfile
	cmd.Stderr = errorfile

	return cmd, nil
}
