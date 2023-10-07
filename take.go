package iter

func (it Iterator) Skip(n int) Iterator {
	return From(func() (v any) {
		for n > 0 {
			if v = it.Next(); v == nil {
				return
			}
			n--
		}
		return it.Next()
	})
}

func (it Iterator) Take(n int) Iterator {
	return From(func() (v any) {
		if n > 0 {
			v = it.Next()
			n--
		}
		return
	})
}

func (it Iterator) Filter(pred func(v any) bool) Iterator {
	return From(func() (v any) {
		v = it.Next()
		for v != nil && !pred(v) && v != nil {
			v = it.Next()
		}
		return
	})
}
