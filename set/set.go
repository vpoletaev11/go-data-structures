package set

type Set struct {
	set map[int]struct{}
}

// Init initializes Set
func Init() Set {
	s := Set{}
	s.set = make(map[int]struct{})
	return s
}

// Len returns lenhgt of Set
func (s *Set) Len() int {
	return len(s.set)
}

// Add adds value in Set
func (s *Set) Add(val int) {
	s.set[val] = struct{}{}
}

// Delete deletes value from Set
func (s *Set) Delete(val int) {
	delete(s.set, val)
}

// Has checks if value exists in Set
func (s *Set) Has(val int) bool {
	_, ok := s.set[val]
	return ok
}

// Clear removes all values from Set
func (s *Set) Clear() {
	s.set = make(map[int]struct{})
}

// GetSet returns all values of Set
func (s *Set) GetSet() []int {
	keys := []int{} //make([]int, len(s.set))
	for k := range s.set {
		keys = append(keys, k)
	}
	return keys
}
