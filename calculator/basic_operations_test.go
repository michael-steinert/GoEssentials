package calculator_test

import (
	"testing"

	"github.com/michael-steinert/GoEssentials/calculator"
	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	// Table Driven Unit Tests
	tests := []struct {
		name        string
		left, right int
		sum         int
	}{
		{"Big positive Integers", 21, 42, 63},
		{"Small positive Integers", 2, 4, 6},
		{"Zeros", 0, 0, 0},
		{"Positive and negative Integers", -21, 42, 21},
	}
	for _, test := range tests {
		// Sub Testing with Test Name
		t.Run(test.name, func(t *testing.T) {
			sum := calculator.Addition(test.left, test.right)
			assert.Equal(t, test.sum, sum)
			// Without Testing Framework / Assertion Library
			if sum != test.sum {
				// Errorf does not interrupt current Test if failed
				t.Errorf("Wrong Result from Addition, got %d, but want %d", sum, test.sum)
				// Fatalf interrupts current Test if failed
				// t.Fatalf("Wrong Result from Addition, got %d, but want %d", sum, exceptedResult)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	// Table Driven Unit Tests
	tests := []struct {
		name                string
		left, right         int
		quotient, remainder int
		withError           bool
	}{
		{"Big positive Integers", 21, 7, 3, 0, false},
		{"Small positive Integers", 5, 2, 2, 1, false},
		{"Zeros", 0, 0, 0, 0, true},
		{"Positive and negative Integers", -21, 7, -3, 0, false},
	}
	for _, test := range tests {
		// Sub Testing with Test Name
		t.Run(test.name, func(t *testing.T) {
			quotient, remainder, err := calculator.Divide(test.left, test.right)
			// Expected Error in Test
			if test.withError {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, test.quotient, quotient)
			assert.Equal(t, test.remainder, remainder)
			assert.NoError(t, err)
		})
	}
}
