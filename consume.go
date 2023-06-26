package iter

func (it Iterator) Last() (v any) {
	if v = it.Next(); v == nil {
		return
	}

	for {
		v1 := it.Next()
		if v1 == nil {
			break
		}
		v = v1
	}
	return
}

func (it Iterator) At(i int) (v any) {
	for i >= 0 {
		if v = it.Next(); v == nil {
			return
		}
		i--
	}
	return
}

func (it Iterator) Reduce(f func(a, b any) any) (v any) {
	if v = it.Next(); v == nil {
		return
	}

	for {
		b := it.Next()
		if b == nil {
			break
		}

		v = f(v, b)
	}
	return
}

func (it Iterator) Fold(st any, f func(st, v any) any) (v any) {
	v = st

	for {
		v1 := it.Next()
		if v1 == nil {
			return
		}

		if v1 = f(v, v1); v1 == nil {
			return
		}
		v = v1
	}
}
