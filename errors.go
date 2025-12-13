package unitor

import "errors"

var (
	ErrParseWrongUnit     = errors.New("input value is a different unit")
	ErrPrefixNotFound     = errors.New("prefix not found")
	ErrUnitless           = errors.New("value is unitless")
	ErrUnitUnknown        = errors.New("unit unknown")
	ErrMissingConversions = errors.New("missing conversions between units")
)
