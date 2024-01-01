package calculator

import "errors"

func Addition(left int, right int) int {
	return left + right
}

func Subtraction(left int, right int) int {
	// var difference int = left - right
	// var difference = left - right
	difference := left - right
	return difference
}

func Divide(left, right int) (quotient int, remainder int, err error) {
	if right == 0 {
		return 0, 0, errors.New("Division by Zero")
	}
	quotient = left / right
	remainder = left % right
	// Naked Return, because all Return Values are assigned with Values
	return
}
