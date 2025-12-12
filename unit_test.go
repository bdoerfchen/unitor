package unitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testUnit = NewUnit("A").WithPrefix("m", 1e-3).Unit()

// Test if Unit.Symbol() returns the units symbol
func TestUnit_Symbol(t *testing.T) {
	assert.Equal(t, "A", testUnit.Symbol())
}

func TestUnit_Value(t *testing.T) {
	v := testUnit.Value(5)
	assert.Equal(t, 5.0, v.Value())
	assert.Equal(t, "A", v.Symbol())
}

func TestUnit_ValuePrefix(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		v, err := testUnit.ValuePrefixed(5, "m")
		assert.NoError(t, err)
		assert.Equal(t, 5.0, v.Value())
		assert.Equal(t, "mA", v.Symbol())
	})

	t.Run("unknown", func(t *testing.T) {
		v, err := testUnit.ValuePrefixed(5, "Gi")
		// Check returns error
		assert.ErrorIs(t, ErrPrefixNotFound, err)

		// Value should be empty
		assert.True(t, v.IsEmpty())
	})
}

func TestUnit_Parse(t *testing.T) {

}
