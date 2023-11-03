package res

type Result[A any] struct {
	v   A
	err error
}

func Wrap[A any](v A, err error) (res Result[A]) {
	return Result[A]{v, err}
}

func Values[A any](res Result[A]) (A, error) {
	return res.v, res.err
}

func Try[A, B any](res Result[A], f func(A) (B, error)) Result[B] {
	if res.err != nil {
		var zero B
		return Result[B]{zero, res.err}
	}
	v, err := f(res.v)
	return Result[B]{v, err}
}
