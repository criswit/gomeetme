package factory


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAvailabilityMap(t *testing.T) {
	actual := CreateAvailabilityMap()
	i := 0
	for _, v := range actual {
		i = i + 1
		assert.Equal(t, v, true)
	}

	assert.Equal(t, i, 86400)
}

