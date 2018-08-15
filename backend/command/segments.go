package command

type CommandSegments []CommandSegment

type CommandSegment struct {
	Key        string
	Base       string
	IsRequired bool
	IsValuable bool
}

func (cs *CommandSegments) ToString() (string, error) {
	return "", nil
}

func (cs *CommandSegments) Load(encoded string) error {
	return nil
}
