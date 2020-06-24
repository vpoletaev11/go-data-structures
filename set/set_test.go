package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/set"
)

var dataset = []int{
	200,
	1,
	14,
	8,
	14,
	16,
	1,
	5,
	9,
	55,
}

func TestSet(t *testing.T) {
	s := set.Init()
	assert.Equal(t, 0, s.Len())

	for _, val := range dataset {
		s.Add(val)
		assert.True(t, s.Has(val))
		s.Delete(val)
		assert.False(t, s.Has(val))
	}

	for _, val := range dataset {
		s.Add(val)
	}
	assert.Equal(t, 8, s.Len())
	assert.ElementsMatch(t, []int{200, 1, 14, 8, 16, 5, 9, 55}, s.GetSet())
	s.Clear()
	assert.Equal(t, []int{}, s.GetSet())

}
