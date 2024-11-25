package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func Generate[T comparable, V any](values []V, transformer func(V) T) Set[T] {
	s := New[T]()
	for _, v := range values {
		s.Add(transformer(v))
	}
	return s
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Equal(s1 Set[T]) bool {
	if len(s) != len(s1) {
		return false
	}
	for k := range s {
		if !s1.Contains(k) {
			return false
		}
	}
	return true
}
