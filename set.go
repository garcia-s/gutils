package gutils

var empty = struct{}{}

type Set[T comparable] struct {
	data map[T]*struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		data: make(map[T]*struct{}),
	}
}

func (s *Set[T]) Exist(id T) bool {
	if s.data[id] != nil {
		return true
	}
	return false
}

func (s *Set[T]) Add(id T) {
	s.data[id] = &empty
}

func (s *Set[T]) ToSlice() []T {
	keys := make([]T, len(s.data))

	i := 0
	for k := range s.data {
		keys[i] = k
		i++
	}
	return keys
}

func (s *Set[T]) Delete(id T) {
	delete(s.data, id)
}
