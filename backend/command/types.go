package command

type CommandType int

const (
	BATCH CommandType = iota
	DAEMON
	INTERACTIVE
)
