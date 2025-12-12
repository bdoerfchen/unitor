package unitor_test

import (
	"testing"

	"github.com/bdoerfchen/unitor"
	"github.com/stretchr/testify/assert"
)

var testUnit = unitor.NewUnit("Test").WithPrefix("k", 1e3).WithPrefix("m", 1e-3).Unit()

func TestValue_Normalize(t *testing.T) {
	value := testUnit.Value(2000).Normalize()

	assert.Equal(t, 2000.0, value.Value())
	assert.Equal(t, "Test", value.Symbol())
}

func TestValue_Reduce(t *testing.T) {
	t.Run("large", func(t *testing.T) {
		value := testUnit.Value(2000).Reduce()
		// Make sure it is reduced
		assert.Equal(t, 2.0, value.Value())
		assert.Equal(t, "kTest", value.Symbol())

		// Make sure reduced is normalized
		assert.Equal(t, 2000.0, value.Normalize().Value())
	})

	t.Run("small", func(t *testing.T) {
		value := testUnit.Value(0.2).Reduce()
		// Make sure it is reduced
		assert.Equal(t, 200.0, value.Value())
		assert.Equal(t, "mTest", value.Symbol())

		// Make sure reduced is normalized
		assert.Equal(t, 0.2, value.Normalize().Value())
	})

	t.Run("base", func(t *testing.T) {
		value := testUnit.Value(5).Reduce()
		// Make sure it is reduced
		assert.Equal(t, 5.0, value.Value())
		assert.Equal(t, "Test", value.Symbol())

		// Make sure reduced is normalized
		assert.Equal(t, 5.0, value.Normalize().Value())
	})
}

func TestValue_In(t *testing.T) {
	t.Run("prefix-only", func(t *testing.T) {
		converted, err := testUnit.Value(5000.0).In("k")

		assert.NoError(t, err)
		assert.Equal(t, 5.0, converted.Value())
		assert.Equal(t, "kTest", converted.Symbol())
	})

	t.Run("full-unit", func(t *testing.T) {
		converted, err := testUnit.Value(5000.0).In("kTest")

		assert.NoError(t, err)
		assert.Equal(t, 5.0, converted.Value())
		assert.Equal(t, "kTest", converted.Symbol())
	})

	t.Run("double", func(t *testing.T) {
		converted, err := testUnit.Value(5000.0).In("k")
		assert.NoError(t, err)
		converted, err = converted.In("m")

		assert.NoError(t, err)
		assert.Equal(t, 5_000_000.0, converted.Value())
		assert.Equal(t, "mTest", converted.Symbol())
	})
}

func TestValue_IsEmpty(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.True(t, unitor.Value{}.IsEmpty())
		unitor.Value{}.Normalize()
	})

	t.Run("not-empty", func(t *testing.T) {
		assert.False(t, unitor.Unitless(10).IsEmpty())
		unitor.Value{}.Normalize()
	})
}

func TestValue_Unitless(t *testing.T) {
	value := unitor.Unitless(10)

	t.Run("isunitless", func(t *testing.T) {
		assert.True(t, value.IsUnitless())
	})

	t.Run("value", func(t *testing.T) {
		assert.Equal(t, 10.0, value.Value())
	})

	t.Run("in", func(t *testing.T) {
		converted, err := value.In("cm")

		assert.Error(t, err)
		assert.True(t, converted.IsUnitless())
		assert.Equal(t, 10.0, converted.Value())
	})

	t.Run("normalize", func(t *testing.T) {
		converted := value.Normalize()

		assert.True(t, converted.IsUnitless())
		assert.Equal(t, 10.0, converted.Value())
	})

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "10", unitor.Unitless(10).String())
		assert.Equal(t, "10.5", unitor.Unitless(10.5).String())
	})
}

func TestValue_String(t *testing.T) {
	t.Run("decimal", func(t *testing.T) {
		assert.Equal(t, "7.2Test", testUnit.Value(7.2).String())
	})

	t.Run("integer", func(t *testing.T) {
		assert.Equal(t, "7Test", testUnit.Value(7).String())
	})

	t.Run("prefixed", func(t *testing.T) {
		value, err := testUnit.Value(7).In("m")
		assert.NoError(t, err)
		assert.Equal(t, "7000mTest", value.String())
	})

	t.Run("custom-format", func(t *testing.T) {
		customFormatUnit := unitor.NewUnit("Test").WithFormat("%.2f", false).Unit()
		assert.Equal(t, "7.20Test", customFormatUnit.Value(7.2).String())
	})

	t.Run("custom-format-front", func(t *testing.T) {
		customFormatUnit := unitor.NewUnit("Test").WithFormat(" %.2f", true).Unit()
		assert.Equal(t, "Test 7.20", customFormatUnit.Value(7.2).String())
	})
}
