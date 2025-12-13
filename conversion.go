package unitor

import (
	"fmt"

	"github.com/dominikbraun/graph"
)

type ConversionFunction func(value float64) float64

type UnitRepository struct {
	units graph.Graph[*Unit, *Unit]
}

type ConversionData struct {
	from      *Unit
	Forwards  ConversionFunction
	Backwards ConversionFunction
}

func ConversionFactor(factor float64) ConversionData {
	return ConversionData{
		Forwards:  func(value float64) float64 { return value * factor },
		Backwards: func(value float64) float64 { return value / factor },
	}
}

func NewRepository() *UnitRepository {
	return &UnitRepository{
		units: graph.New(func(u *Unit) *Unit { return u }),
	}
}

func (r *UnitRepository) AddConversion(from *Unit, to *Unit, data ConversionData) *UnitRepository {
	_ = r.units.AddVertex(from)
	_ = r.units.AddVertex(to)
	data.from = from
	err := r.units.AddEdge(from, to, graph.EdgeData(data))
	if err != nil {
		panic("Unable to store conversion")
	}

	return r
}

func (r *UnitRepository) Convert(from Value, to *Unit) (Value, error) {
	if from.IsUnitless() {
		return Value{}, ErrUnitless
	}

	// Return given value directly when trying to convert to same unit
	if from.unitPrefix.unit == to {
		return from, nil
	}

	// Get path
	path, err := graph.ShortestPath(r.units, from.unitPrefix.unit, to)
	if err != nil {
		return from, ErrMissingConversions
	}

	// Convert between paths
	var conversionFunc ConversionFunction
	var currentUnit *Unit
	var currentValue Value
	for i, next := range path {
		if i == 0 {
			currentValue = from.Normalize()
			currentUnit = from.unitPrefix.unit
			continue
		}

		edge, err := r.units.Edge(currentUnit, next)
		if err != nil {
			return from, fmt.Errorf("expected edge on conversion path, but received error: %w", err)
		}
		conversionInfo, ok := edge.Properties.Data.(ConversionData)
		if !ok {
			return from, fmt.Errorf("conversion data missing: %w", err)
		}

		if conversionInfo.from == currentUnit {
			conversionFunc = conversionInfo.Forwards
		} else {
			conversionFunc = conversionInfo.Backwards
		}

		// Finish conversion
		currentUnit = next
		currentValue = next.Value(conversionFunc(currentValue.value)) //Already normalized
	}

	return currentValue, nil
}
