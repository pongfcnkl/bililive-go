package recorders

import "errors"

var (
	ErrRecorderExist    = errors.New("recorder is exist")
	ErrRecorderNotExist = errors.New("recorder is not exist")
)
