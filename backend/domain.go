package entrance

import "time"

type CommandType int

const (
	BATCH CommandType = iota
	DAEMON
	INTERACTIVE
)

type CommandSegment struct {
	Key        string
	Base       string
	IsRequired bool
	IsValuable bool
}

type Base struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Command struct {
	Base
	Name            string
	CommandType     CommandType
	CommandSegments []CommandSegment
}

type JobStatus int

const (
	WAITING JobStatus = iota //0
	PREPARE                  //1
	RUNNING                  //2
	DONE                     //3
	STOP                     //4
)

type Arguments map[string]string

type Job struct {
	Base
	Status    JobStatus
	CommandID uint
	Arguments *Arguments
	SysCmd    string
}

type QueryCondition map[string]interface{}

type ExecuteResult struct {
	OutputData *[]string
	ErrorData  *[]string
}
