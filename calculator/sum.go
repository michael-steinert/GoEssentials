package calculator

// For Loop
func Sum(min, max int) int {
	sum := 0
	for i := min; i < max; i++ {
		sum += i
	}
	return sum
}

// While Loop
func SumUntil(max int) int {
	sum := 1
	for sum < max {
		sum += sum
	}
	return sum
}
