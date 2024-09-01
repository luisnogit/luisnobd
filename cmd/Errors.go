package main

import "errors"

var (
	ErrInvalidMeta = errors.New("invalid meta command"	)
	ErrUnrecognizedStm = errors.New("unknown statement")
	ErrInvalidSyntax = errors.New("invalid syntax")
	
)