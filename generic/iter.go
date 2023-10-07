package iter

const (
	Continue = true
	Break    = false
)

type Iterator[T any] interface {
	Next() (T, bool)
}

type IteratorFunc[T any] func() (T, bool)

func (it IteratorFunc[T]) Next() (T, bool) {
	return it()
}

func Iter[T any](sli []T) Iterator[T] {
	i := 0
	return IteratorFunc[T](func() (T, bool) {
		if i < len(sli) {
			i0 := i
			i++
			return sli[i0], true
		}
		var zero T
		return zero, false
	})
}

func Range(a, b int, step int) Iterator[int] {
	i := a
	return IteratorFunc[int](func() (int, bool) {
		if i < b {
			i0 := i
			i += step
			return i0, true
		}
		return 0, false
	})
}

func Empty[T any]() Iterator[T] {
	return IteratorFunc[T](func() (T, bool) {
		var zero T
		return zero, false
	})
}

func Repeat[T any](t T) Iterator[T] {
	return IteratorFunc[T](func() (T, bool) {
		return t, true
	})
}

func Chain[T any](it1, it2 Iterator[T]) Iterator[T] {
	return IteratorFunc[T](func() (T, bool) {
		t, ok := it1.Next()
		if ok {
			return t, true
		}
		return it2.Next()
	})
}

type zip[A, B any] struct {
	a A
	b B
}

func Zip[A, B any](itA Iterator[A], itB Iterator[B]) Iterator[zip[A, B]] {
	return IteratorFunc[zip[A, B]](func() (zip[A, B], bool) {
		a, okA := itA.Next()
		b, okB := itB.Next()
		if !okA || !okB {
			return zip[A, B]{}, false
		}
		return zip[A, B]{a, b}, true
	})
}
