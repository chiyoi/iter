package res

func Try[A, B any](a A, err error, f func(A) (B, error)) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f(a)
}
