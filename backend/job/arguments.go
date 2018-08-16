package job

import "encoding/json"

type Arguments map[string]string

func NewArguments(str string) (*Arguments, error) {
	var argument = make(Arguments)
	if err := argument.Load(str); err != nil {
		return nil, err
	}
	return &argument, nil
}

func (arguments *Arguments) ToString() (string, error) {
	bytes, err := json.Marshal(*arguments)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (arguments *Arguments) Load(str string) error {
	return json.Unmarshal([]byte(str), &arguments)
}

func (arguments *Arguments) Get(key string) (string, bool) {
	val, ok := (*arguments)[key]
	return val, ok
}
