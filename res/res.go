package res

func Then[B, A any](a A, err error, f func(A) (B, error)) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f(a)
}

func And[A any](err error, f func() (A, error)) (A, error) {
	if err != nil {
		var zero A
		return zero, err
	}
	return f()
}
