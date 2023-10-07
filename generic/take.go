package iter

func Take[T any](it Iterator[T], count int) Iterator[T] {
	i := 0
	return IteratorFunc[T](func() (t T, ok bool) {
		if i < count {
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
