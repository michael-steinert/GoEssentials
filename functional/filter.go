package functional

// Generics
func Filter[T any](values []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, value := range values {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
