package unitor

import (
	"fmt"
	"strings"
)

type Unit struct {
	symbol   string
	prefixes *prefixManager

	symbolBeforeValue bool
	valueFormat       string
}

func (u *Unit) basePrefix() *unitPrefix {
	// Find or add base prefix
	prefix := u.prefixes.For("")
	if prefix == nil {
		prefix = &unitPrefix{
			unit:       u,
			BaseFactor: 1,
		}
		_ = u.prefixes.Add(prefix)
	}

	return prefix
}

func (u *Unit) Symbol() string {
	return u.symbol
}

func (u *Unit) Parse(input string) (Value, error) {
	value, unitText, err := DefaultParser.Parse(input, u.symbol)
	if err != nil {
		return Value{}, fmt.Errorf("can not parse to unit: %w", err)
	}

	// If no unit, return unitless value
	if unitText == "" {
		return Value{value: value}, nil
	}

	// If no prefix, return value directly
	if unitText == u.symbol {
		return Value{
			value:      value,
			unitPrefix: u.basePrefix(),
		}, nil
	}

	// Check if of symbol
	prefixText, endsInUnitSymbol := strings.CutSuffix(unitText, u.symbol)
	if !endsInUnitSymbol {
		return Value{}, ErrParseWrongUnit
	}

	return u.ValuePrefixed(value, prefixText)
}

func (u *Unit) Value(value float64) Value {
	return Value{
		value:      value,
		unitPrefix: u.basePrefix(),
	}
}

func (u *Unit) ValuePrefixed(value float64, prefixText string) (Value, error) {
	// Remove unit symbol if provided in prefix
	prefixText, _ = strings.CutSuffix(prefixText, u.symbol)

	prefix := u.prefixes.For(prefixText)
	if prefix == nil {
		return Value{}, ErrPrefixNotFound
	}

	return Value{
		value:      value,
		unitPrefix: prefix,
	}, nil
}
