package entrance

type App interface {
	StorageFeature
	ExecuteFeature
}

type StorageFeature interface {
	CreateCommand(name string, commandtype CommandType, commandSegments []CommandSegment) (*Command, error)
	Command(cID uint) (*Command, error)
	Commands(*QueryCondition) (*[]Command, error)
	SaveCommand(*Command) error
	CreateJob(uint, *Arguments) (*Job, error)
	Job(jID uint) (*Job, error)
	Jobs(*QueryCondition) (*[]Job, error)
	SaveJob(*Job) error
}

type ExecuteFeature interface {
	Execute(*Job) error
	GetResult(*Job) (*ExecuteResult, error)
}
