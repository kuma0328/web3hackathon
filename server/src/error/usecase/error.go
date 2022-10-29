package usecase_error

import "errors"

var (
	IdEmptyError       = errors.New("id empty")
	NameEmptyError     = errors.New("name empty")
	PassWordEmptyError = errors.New("password empty")
)
