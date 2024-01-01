package functional_test

import (
	"testing"

	"github.com/michael-steinert/GoEssentials/functional"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		[]int{},
		functional.Map([]int{}, func(prime int) int {
			return prime * prime
		}),
	)
	assert.Equal(t,
		[]int{4, 9, 25},
		functional.Map([]int{2, 3, 5}, func(prime int) int {
			return prime * prime
		}),
	)
	assert.Equal(t,
		[]int{5, 6},
		functional.Map([]string{"Hello", "Bruno!"}, func(word string) int {
			return len(word)
		}),
	)
}
