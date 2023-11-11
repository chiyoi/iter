package opt

func Or[A comparable](a A, f func() A) A {
	var zero A
	if a != zero {
		return a
	}
	return f()
}
