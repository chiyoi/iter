package iter

func Map[A, B any](it Iterator[A], f func(A) B) Iterator[B] {
	return &mapped[A, B]{it, f}
}

func Scan[T, St any](it Iterator[T], st St, f func(St, T) St) Iterator[St] {
	return &scanned[T, St]{it, st, f}
}

type mapped[A, B any] struct {
	it Iterator[A]
	f  func(A) B
}

func (m *mapped[A, B]) Next() (b B, ok bool) {
	a, ok := m.it.Next()
	if !ok {
		return
	}
	return m.f(a), true
}

type scanned[A, St any] struct {
	it Iterator[A]
	st St
	f  func(St, A) St
}

func (s *scanned[A, St]) Next() (st St, ok bool) {
	a, ok := s.it.Next()
	if !ok {
		return
	}
	st = s.f(s.st, a)
	s.st = st
	return
}
