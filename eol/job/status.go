package job

type JobStatus int

const (
	WAITING JobStatus = iota //0
	PREPARE                  //1
	RUNNING                  //2
	DONE                     //3
	STOP                     //4
)
