package functional_test

import (
	"testing"

	"github.com/michael-steinert/GoEssentials/functional"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert.Equal(t,
		[]int{},
		functional.Filter([]int{}, func(value int) bool {
			return value < 4
		}),
	)
	assert.Equal(t,
		[]int{2, 3},
		functional.Filter([]int{2, 3, 5}, func(value int) bool {
			return value < 4
		}),
	)
}
