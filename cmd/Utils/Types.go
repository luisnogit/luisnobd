package utils

import "errors"

type StatementType uint8
const (
	INSERT = iota
	SELECT
)

type Statement struct {
	Row_toinsert Row
	Tp           StatementType
	Input        string
}
type Row struct {
	Id       int
	Username string
	Email    string
}

var (
	ErrInvalidMeta     = errors.New("invalid meta command")
	ErrUnrecognizedStm = errors.New("unknown statement")
	ErrInvalidSyntax   = errors.New("invalid syntax")
)
