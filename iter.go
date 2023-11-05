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
	return IteratorFunc[T](func() (v T, ok bool) {
		if i < len(sli) {
			v, ok = sli[i], true
			i++
			return
		}
		var zero T
		return zero, false
	})
}

func Keys[K comparable, A any](m map[K]A) Iterator[K] {
	keys := make([]K, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return Iter(keys)
}

func IterRev[T any](sli []T) Iterator[T] {
	i := len(sli) - 1
	return IteratorFunc[T](func() (v T, ok bool) {
		if i >= 0 {
			v, ok = sli[i], true
			i--
			return
		}
		var zero T
		return zero, false
	})
}

func Range(a, b int, step int) Iterator[int] {
	if step == 0 || a <= b && step < 0 || a >= b && step > 0 {
		return Empty[int]()
	}
	i := a
	return IteratorFunc[int](func() (int, bool) {
		if step > 0 && i >= b || step < 0 && i <= b {
			return 0, false
		}
		i0 := i
		i += step
		return i0, true
	})
}

func Empty[T any]() Iterator[T] {
	return IteratorFunc[T](func() (T, bool) {
		var zero T
		return zero, false
	})
}

func Once[T any](v T) Iterator[T] {
	var flag bool
	return IteratorFunc[T](func() (T, bool) {
		if flag {
			var zero T
			return zero, false
		}
		flag = true
		return v, true
	})
}

func Repeat[T any](v T) Iterator[T] {
	return IteratorFunc[T](func() (T, bool) {
		return v, true
	})
}

func Chain[T any](itA, itB Iterator[T]) Iterator[T] {
	return IteratorFunc[T](func() (T, bool) {
		t, ok := itA.Next()
		if ok {
			return t, true
		}
		return itB.Next()
	})
}

func Zip[A, B any](itA Iterator[A], itB Iterator[B]) Iterator[ZipItem[A, B]] {
	return IteratorFunc[ZipItem[A, B]](func() (ZipItem[A, B], bool) {
		a, okA := itA.Next()
		b, okB := itB.Next()
		if !okA || !okB {
			return ZipItem[A, B]{}, false
		}
		return ZipItem[A, B]{a, b}, true
	})
}

func Map[A, B any](it Iterator[A], f func(A) B) Iterator[B] {
	return IteratorFunc[B](func() (b B, ok bool) {
		a, ok := it.Next()
		if !ok {
			return
		}
		return f(a), true
	})
}

func Scan[T, St any](it Iterator[T], st St, f func(St, T) (St, bool)) Iterator[St] {
	return IteratorFunc[St](func() (out St, ok bool) {
		v, ok := it.Next()
		if !ok {
			return
		}
		st1, ok := f(st, v)
		if !ok {
			return st, false
		}
		st = st1
		return st, true
	})
}

func Take[T any](it Iterator[T], count int) Iterator[T] {
	i := 0
	return IteratorFunc[T](func() (v T, ok bool) {
		if i < count {
			i++
			return it.Next()
		}
		return
	})
}

func Skip[T any](it Iterator[T], count int) Iterator[T] {
	var flag bool
	return IteratorFunc[T](func() (v T, ok bool) {
		if !flag {
			for i := 0; i < count; i++ {
				_, ok = it.Next()
				if !ok {
					return
				}
			}
			flag = true
		}
		return it.Next()
	})
}

func Filter[T any](it Iterator[T], f func(T) bool) Iterator[T] {
	return IteratorFunc[T](func() (v T, ok bool) {
		v, ok = it.Next()
		if !ok {
			return
		}
		for !f(v) {
			v, ok = it.Next()
			if !ok {
				return
			}
		}
		return
	})
}

func Last[T any](it Iterator[T]) (v T, ok bool) {
	return Reduce(it, func(_, t T) T {
		return t
	})
}

func At[T any](it Iterator[T], i int) (v T, ok bool) {
	if i < 0 {
		return
	}
	return Skip(it, i).Next()
}

func Reduce[T any](it Iterator[T], f func(T, T) T) (ans T, ok bool) {
	ans, ok = it.Next()
	if !ok {
		return
	}
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		ans = f(ans, v)
	}
	return
}

func Fold[T, St any](it Iterator[T], st St, f func(St, T) (St, bool)) St {
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		st1, ok := f(st, v)
		if !ok {
			return st
		}
		st = st1
	}
	return st
}

func Collect[T any](it Iterator[T]) (ans []T) {
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		ans = append(ans, v)
	}
	return
}

func Count[T any](it Iterator[T]) (count int) {
	for {
		_, ok := it.Next()
		if !ok {
			break
		}
		count++
	}
	return
}

type ZipItem[A, B any] struct {
	A A
	B B
}
