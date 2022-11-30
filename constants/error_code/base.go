package error_code

import "errors"

var (
	ErrAlreadyExists = errors.New("record already exists")
)
