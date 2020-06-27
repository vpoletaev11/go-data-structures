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

func BenchmarkSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := set.Init()
		b.StartTimer()

		for _, val := range dataset {
			s.Add(val)
		}
	}
}

func BenchmarkSetDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := set.Init()
		for _, val := range dataset {
			s.Add(val)
		}
		b.StartTimer()

		for _, val := range dataset {
			s.Delete(val)
		}
	}
}

func BenchmarkSetHas(b *testing.B) {
	s := set.Init()
	for _, val := range dataset {
		s.Add(val)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, val := range dataset {
			s.Has(val)
		}
	}
}

func BenchmarkSetClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := set.Init()
		for _, val := range dataset {
			s.Add(val)
		}
		b.StartTimer()

		s.Clear()
	}
}

func BenchmarkGetSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := set.Init()
		for _, val := range dataset {
			s.Add(val)
		}
		b.StartTimer()

		s.GetSet()
	}
}
