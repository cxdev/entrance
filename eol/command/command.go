package command

import "entrance/eol/platform"

type Command struct {
	platform.BaseEntity
	Name            string
	CommandType     CommandType
	CommandSegments CommandSegments
}

func New(name string, commandType CommandType, segments string) (*Command, error) {
	var cs CommandSegments
	if err := cs.Load(segments); err != nil {
		return nil, err
	}

	return &Command{platform.BaseEntity{}, name, commandType, cs}, nil
}
