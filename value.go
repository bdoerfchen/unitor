package unitor

import (
	"fmt"
)

type Value struct {
	value      float64
	unitPrefix *unitPrefix
}

func Unitless(value float64) Value {
	return Value{
		value: value,
	}
}

func (v Value) IsEmpty() bool {
	return v.IsUnitless() && v.value == 0
}

func (v Value) IsUnitless() bool {
	return v.unitPrefix == nil
}

func (v Value) String() string {
	if v.IsUnitless() {
		return fmt.Sprint(v.value)
	}

	if v.unitPrefix.unit.symbolBeforeValue {
		return fmt.Sprintf("%s"+v.unitPrefix.unit.valueFormat, v.Symbol(), v.value)
	} else {
		return fmt.Sprintf(v.unitPrefix.unit.valueFormat+"%s", v.value, v.Symbol())
	}
}

func (v Value) Symbol() string {
	if v.IsUnitless() {
		return ""
	}

	return v.unitPrefix.String()
}

// Reduce value to closest prefix
func (v Value) Reduce() Value {
	if v.IsUnitless() {
		return v
	}

	base := v.Normalize()
	newPrefix := v.unitPrefix.unit.prefixes.Closest(base.value)
	return Value{
		unitPrefix: newPrefix,
		value:      base.value / newPrefix.BaseFactor,
	}
}

// Normalizes value for unit
func (v Value) Normalize() Value {
	if v.IsUnitless() {
		return v
	}

	return Value{
		value:      v.unitPrefix.BaseFactor * v.value,
		unitPrefix: v.unitPrefix.unit.basePrefix(),
	}
}

func (v Value) In(prefix string) (Value, error) {
	if v.IsUnitless() {
		return v, ErrUnitless
	}

	// Value with prefix, but value not converted
	converted, err := v.unitPrefix.unit.ValuePrefixed(v.value, prefix)
	if err != nil {
		return v, err
	}

	// Calculate value for new prefix
	converted.value *= v.unitPrefix.BaseFactor / converted.unitPrefix.BaseFactor

	return converted, nil
}

func (v Value) Value() float64 {
	return v.value
}
