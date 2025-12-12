package unitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnitSymbol(t *testing.T) {
	unit := NewUnit("A").Unit()
	assert.Equal(t, "A", unit.Symbol())
}
