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

func Range(a, b int, step int) Iterator[int] {
	i := a
	return IteratorFunc[int](func() (int, bool) {
		if i < b {
			i0 := i
			i++
			return i0, true
		}
		return 0, false
	})
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
