package iter

func (it Iterator) Map(f func(a any) (b any)) Iterator {
	return From(func() (v any) {
		v0 := it.Next()
		if v0 != nil {
			v = f(v0)
		}
		return
	})
}

func (it Iterator) Scan(st any, f func(st any, v any) any) Iterator {
	return From(func() (v any) {
		v0 := it.Next()
		if v0 != nil {
			st = f(st, v0)
			v = st
		}
		return
	})
}
