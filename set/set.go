package set

// Set unsorted set
type Set[T comparable] map[T]struct{}

func New[T comparable](hint int) Set[T] {
	return make(Set[T], hint)
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Del(v T) {
	delete(s, v)
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}
