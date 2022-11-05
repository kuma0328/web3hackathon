package usecase

import "errors"

var (
	IdEmptyError       = errors.New("id empty")
	NameEmptyError     = errors.New("name empty")
	MailEmptyError     = errors.New("mail empty")
	PassWordEmptyError = errors.New("password empty")
)
