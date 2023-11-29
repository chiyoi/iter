package opt

// A (Address)
func A[A any](a A) *A {
	return &a
}

func Or[A comparable](x, y A) A {
	var zero A
	if x == zero {
		return y
	}
	return x
}
