package unitor_test

import (
	"testing"

	"github.com/bdoerfchen/unitor"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Conversion(t *testing.T) {
	// Setup units A, B, C
	// with conversions A --*2--> B --*4--> C
	unitA := unitor.NewUnit("A").Unit()
	unitB := unitor.NewUnit("B").Unit()
	unitC := unitor.NewUnit("C").Unit()
	repository := unitor.NewRepository().
		AddConversion(unitA, unitB, unitor.ConversionFactor(2)).
		AddConversion(unitB, unitC, unitor.ConversionFactor(4))

	t.Run("forward", func(t *testing.T) {
		valueA := unitA.Value(5)
		valueC, err := repository.Convert(valueA, unitC)

		assert.NoError(t, err)
		assert.Equal(t, 40.0, valueC.Value())
		assert.Equal(t, "C", valueC.Symbol())
	})

	t.Run("backwards", func(t *testing.T) {
		valueC := unitC.Value(40)
		valueA, err := repository.Convert(valueC, unitA)

		assert.NoError(t, err)
		assert.Equal(t, 5.0, valueA.Value())
		assert.Equal(t, "A", valueA.Symbol())
	})

	t.Run("same", func(t *testing.T) {
		valueA := unitA.Value(5)
		valueA2, err := repository.Convert(valueA, unitA)

		assert.NoError(t, err)
		assert.Equal(t, 5.0, valueA2.Value())
		assert.Equal(t, "A", valueA2.Symbol())
	})

	t.Run("unknown", func(t *testing.T) {
		// New unit not in repository
		unitD := unitor.NewUnit("D").Unit()

		// Try convert A->D
		valueA := unitA.Value(2)
		_, err := repository.Convert(valueA, unitD)

		assert.Error(t, err)
	})
}
