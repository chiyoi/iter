package res

// R (Redeem)
func R[A, B any](a A, err error, f func(A) (B, error)) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f(a)
}

// C (Consume)
func C[A any](a A, err error, f func(A) error) error {
	if err != nil {
		return err
	}
	return f(a)
}

// M (Map)
func M[A, B any](a A, err error, f func(A) B) (B, error) {
	if err != nil {
		var zero B
		return zero, err
	}
	return f(a), nil
}

type Hook[A any] func(A) (A, error)

func ComposedHooks[A any](hooks ...Hook[A]) Hook[A] {
	return func(a A) (A, error) {
		var err error
		for _, hook := range hooks {
			if hook != nil {
				a, err = R(a, err, hook)
			}
		}
		return a, err
	}
}
