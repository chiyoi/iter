package iter

import "cmp"

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

func Scan[T, St any](it Iterator[T], st St, f func(St, T) St) Iterator[St] {
	return IteratorFunc[St](func() (st St, ok bool) {
		a, ok := it.Next()
		if !ok {
			return
		}
		st = f(st, a)
		return st, true
	})
}

func Take[T any](it Iterator[T], count int) Iterator[T] {
	i := 0
	return IteratorFunc[T](func() (t T, ok bool) {
		if i < count {
			i++
			return it.Next()
		}
		return
	})
}

func Skip[T any](it Iterator[T], count int) Iterator[T] {
	var flag bool
	return IteratorFunc[T](func() (t T, ok bool) {
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
	return IteratorFunc[T](func() (t T, ok bool) {
		t, ok = it.Next()
		if !ok {
			return
		}
		for !f(t) {
			t, ok = it.Next()
			if !ok {
				return
			}
		}
		return
	})
}

func Last[T any](it Iterator[T]) (t T, ok bool) {
	return Reduce(it, func(_, t T) T {
		return t
	})
}

func At[T any](it Iterator[T], i int) (t T, ok bool) {
	return Skip(it, i).Next()
}

func Max[T cmp.Ordered](it Iterator[T]) (t T, ok bool) {
	return Reduce(it, func(a, b T) T {
		if a > b {
			return a
		}
		return b
	})
}

func Min[T cmp.Ordered](it Iterator[T]) (t T, ok bool) {
	return Reduce(it, func(a, b T) T {
		if a < b {
			return a
		}
		return b
	})
}

func Sum[T cmp.Ordered](it Iterator[T]) (t T, ok bool) {
	return Reduce(it, func(a, b T) T {
		return a + b
	})
}

func Reduce[T any](it Iterator[T], f func(T, T) T) (ans T, ok bool) {
	ans, ok = it.Next()
	if !ok {
		return
	}
	for {
		t, ok := it.Next()
		if !ok {
			break
		}
		ans = f(ans, t)
	}
	return
}

func Fold[T, St any](it Iterator[T], st St, f func(St, T) St) St {
	for {
		t, ok := it.Next()
		if !ok {
			break
		}
		st = f(st, t)
	}
	return st
}

func Collect[T any](it Iterator[T]) (ans []T) {
	for {
		t, ok := it.Next()
		if !ok {
			break
		}
		ans = append(ans, t)
	}
	return
}

type ZipItem[A, B any] struct {
	a A
	b B
}
