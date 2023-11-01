package iter

func Call[A, B any](f func(A) B, x A) func() B {
	return func() B {
		return f(x)
	}
}

func Y[A, B any](f func(func(A) B) func(A) B) func(A) B {
	return f(func(x A) B {
		return Y(f)(x)
	})
}
