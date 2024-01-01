package functional

// Generics
func Reduce[T, U any](values []T, f func(U, T) U, initialValue U) U {
	result := initialValue
	for _, value := range values {
		result = f(result, value)
	}
	return result
}
