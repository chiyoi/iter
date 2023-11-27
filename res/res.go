package res

func Then[B, A any](a A, err error, f func(A) (B, error)) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f(a)
}

type None = struct{}

type Hook[T any] func(T) (T, error)

func ComposedHooks[T any](hooks ...Hook[T]) Hook[T] {
	return func(t T) (T, error) {
		var err error
		for _, hook := range hooks {
			t, err = Then(t, err, hook)
		}
		return t, err
	}
}
