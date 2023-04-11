package calculator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"belajar-go/calculator"
)

func TestAddition(t *testing.T) {
	result := calculator.Addition(2, 3)
	assert.Equal(t, 5, result)
}

func TestSubtraction(t *testing.T) {
	result := calculator.Subtraction(5, 3)
	assert.Equal(t, 2, result)
}

func TestMultiplication(t *testing.T) {
	result := calculator.Multiplication(2, 3)
	assert.Equal(t, 6, result)
}

func TestDivision(t *testing.T) {
	t.Run("division by non-zero number", func(t *testing.T) {
		result := calculator.Division(6, 2)
		assert.Equal(t, 3, result)
	})

	t.Run("division by zero", func(t *testing.T) {
		assert.PanicsWithValue(t, "division by zero", func() {
			calculator.Division(6, 0)
		})
	})
}
