package calculator

import (
	"errors"
	"math"
)

func Abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func SquareRoot(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Value can not be negative")
	}
	return math.Sqrt(value), nil
}

func IsSquareNumber(value int) bool {
	squareRoot := math.Sqrt(float64(value))
	return squareRoot == math.Floor(squareRoot)
}

func RunOperation(operation string, left, right int) int {
	var result int
	switch operation {
	case "add":
		result = left + right
	case "subtract":
		result = left - right
	}
	return result
}

func SumFromLeftToRight(left, right int) int {
	return processFromLeftToRight(left, right, 0, func(x, y int) int {
		return x + y
	})
}

func MultiplyFromLeftToRight(left, right int) int {
	return processFromLeftToRight(left, right, 1, func(x, y int) int {
		return x * y
	})
}

// Higher Order Function
func processFromLeftToRight(left, right, initValue int, fn func(int, int) int) int {
	increment := increment(1)
	if left > right {
		return initValue
	}
	return fn(left, processFromLeftToRight(increment(left), right, initValue, fn))
}

func increment(left int) func(int) int {
	return func(increment int) int {
		return left + increment
	}
}
