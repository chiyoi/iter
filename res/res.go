package res

func Then[B, A any](a A, err error, f func(A) (B, error)) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f(a)
}

func And[B any](err error, f func() (B, error)) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f()
}
