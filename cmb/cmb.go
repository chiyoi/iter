package cmb

func Call[A, B any](f func(A) B, a A) func() B {
	return func() B {
		return S(K[A](f), I[A], a)
	}
}

func Literal[A any](a A) func() A {
	return func() A {
		return I(a)
	}
}

func I[X any](x X) X {
	return x
}

func K[Y, X any](x X) func(Y) X {
	return func(Y) X {
		return x
	}
}

func S[A, B, Z any](x func(Z) func(A) B, y func(Z) A, z Z) B {
	return x(z)(y(z))
}

func Y[A, B any](f func(func(A) B) func(A) B) func(A) B {
	return f(func(x A) B {
		return Y(f)(x)
	})
}
