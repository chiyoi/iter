package cmb

func Y[A, B any](f func(func(A) B) func(A) B) func(A) B {
	return f(func(x A) B {
		return Y(f)(x)
	})
}
