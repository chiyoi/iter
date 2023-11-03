package result

type result[A any] struct {
	v   A
	err error
}

func Result[A any](v A, err error) (res result[A]) {
	return result[A]{v, err}
}

func Values[A any](res result[A]) (A, error) {
	return res.v, res.err
}

func Try[A, B any](res result[A], f func(A) (B, error)) result[B] {
	if res.err != nil {
		var zero B
		return result[B]{zero, res.err}
	}
	v, err := f(res.v)
	return result[B]{v, err}
}
