package opt

func Then[A, B comparable](a A, f func(A) B) B {
	var zero A
	if a == zero {
		var zero B
		return zero
	}
	return f(a)
}

func Or[A comparable](a A, f func() A) A {
	var zero A
	if a != zero {
		return a
	}
	return f()
}
