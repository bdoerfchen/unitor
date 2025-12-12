package unitor

import "errors"

var (
	ErrParseWrongUnit = errors.New("input value is a different unit")
	ErrPrefixNotFound = errors.New("prefix not found")
)
