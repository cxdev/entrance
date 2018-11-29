package command

import "encoding/json"

type CommandSegment struct {
	Key        string
	Base       string
	IsRequired bool
	IsValuable bool
}

type CommandSegments []CommandSegment

func NewCommandSegments(str string) (*CommandSegments, error) {
	segments := make(CommandSegments, 0)
	if err := segments.Load(str); err != nil {
		return nil, err
	}
	return &segments, nil
}

func (segments *CommandSegments) ToString() (string, error) {
	bytes, err := json.Marshal(*segments)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (segments *CommandSegments) Load(str string) error {
	return json.Unmarshal([]byte(str), &segments)
}
