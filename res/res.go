package res

func R[B, A any](a A, err error, f func(A) (B, error)) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f(a)
}

func C[T any](v T, err error, f func(T) error) error {
	if err != nil {
		return err
	}
	return f(v)
}

type Hook[T any] func(T) (T, error)

func ComposedHooks[T any](hooks ...Hook[T]) Hook[T] {
	return func(t T) (T, error) {
		var err error
		for _, hook := range hooks {
			t, err = R(t, err, hook)
		}
		return t, err
	}
}
