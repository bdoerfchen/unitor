package unitor

type UnitBuilder struct {
	unit *Unit
}

// Creates a new unit from a builder.
// Its methods are chainable. Errors are resulting in panics.
func NewUnit(symbol string) *UnitBuilder {
	unit := &Unit{
		symbol:            symbol,
		prefixes:          NewPrefixes(),
		valueFormat:       "%v",
		symbolBeforeValue: false,
	}

	return &UnitBuilder{unit: unit}
}

func (b *UnitBuilder) Unit() *Unit {
	return b.unit
}

func (b *UnitBuilder) WithFormat(valueFormat string, unitSymbolBeforeValue bool) *UnitBuilder {
	b.unit.valueFormat = valueFormat
	b.unit.symbolBeforeValue = unitSymbolBeforeValue

	return b
}

func (b *UnitBuilder) WithPrefix(prefix string, factor float64, aliases ...string) *UnitBuilder {
	if b == nil {
		b = NewUnit("")
	}

	// Register prefix for aliases
	basePrefix := &unitPrefix{
		unit:       b.unit,
		BaseFactor: factor,
		Prefix:     prefix,
	}
	err := b.unit.prefixes.Add(basePrefix, aliases...)
	if err != nil {
		panic(err)
	}

	return b
}

func (b *UnitBuilder) WithImportedPrefixes(base *Unit) *UnitBuilder {
	// Add main prefixes
	for _, mainPrefix := range base.prefixes.mainSorted {
		if mainPrefix == nil {
			continue
		}

		b.unit.prefixes.Add(
			&unitPrefix{
				unit:       b.unit,
				Prefix:     mainPrefix.Prefix,
				BaseFactor: mainPrefix.BaseFactor,
			},
			mainPrefix.Prefix,
		)
	}

	// Add aliases
	for text, prefix := range base.prefixes.all {
		// Skip main prefix
		if text == prefix.Prefix {
			continue
		}

		b.unit.prefixes.AddAlias(prefix.Prefix, text)
	}

	return b
}

func (b *UnitBuilder) WithPrefixAlias(alias string, forPrefix string) *UnitBuilder {
	if err := b.unit.prefixes.AddAlias(forPrefix, alias); err != nil {
		panic(err)
	}

	return b
}
