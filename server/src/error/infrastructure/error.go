package infra

import "errors"

var (
	DbConnectError  = errors.New("faild to connect db server")
	QueryError      = errors.New("failed to db query")
	ExecError       = errors.New("failed to sql exec")
	LastIdError     = errors.New("failed to get last insert id")
	RowsScanError   = errors.New("failed to rows scan")
	StatementError  = errors.New("failed to statement")
	GetArticleError = errors.New("failed to getarticle")
)
