package iter

type C[A, B any] func(C[A, B]) func(A) B

func Y[A, B any](f func(func(A) B) func(A) B) func(A) B {
	return (func(c C[A, B]) func(A) B {
		return f(func(a A) B {
			return c(c)(a)
		})
	})(func(c C[A, B]) func(A) B {
		return f(func(a A) B {
			return c(c)(a)
		})
	})
}
