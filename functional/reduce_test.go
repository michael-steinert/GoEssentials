package functional_test

import (
	"testing"

	"github.com/michael-steinert/GoEssentials/functional"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	assert.Equal(t,
		0,
		functional.Reduce([]int{}, func(sum int, value int) int {
			return sum + value
		}, 0),
	)
	assert.Equal(t,
		6,
		functional.Reduce([]int{1, 2, 3}, func(sum int, value int) int {
			return sum + value
		}, 0),
	)
	assert.Equal(t,
		6,
		functional.Reduce([]int{1, 2, 3}, func(product int, value int) int {
			return product * value
		}, 1),
	)
}
