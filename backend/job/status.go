package job

type JobStatus int

const (
	WAITING JobStatus = iota
	PREPARE
	RUNNING
	DONE
	STOP
)
