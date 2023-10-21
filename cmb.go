package iter

func I[A any](x A) A {
	return x
}

func Y[A, B any](f func(func(A) B) func(A) B) func(A) B {
	return f(func(x A) B {
		return Y(f)(x)
	})
}

func S[A, B, C any](x func(A) func(B) C) func(y func(A) B) func(z A) C {
	return func(y func(A) B) func(z A) C {
		return func(z A) C {
			return x(z)(y(z))
		}
	}
}

func B[A, B, C any](x func(B) C) func(y func(A) B) func(z A) C {
	return func(y func(A) B) func(z A) C {
		return func(z A) C {
			return x(y(z))
		}
	}
}

func C[A, B, C any](x func(A) func(B) func(A) C) func(y B) func(z A) C {
	return func(y B) func(z A) C {
		return func(z A) C {
			return x(z)(y)(z)
		}
	}
}

func W[A, B any](x func(A) func(A) B) func(y A) B {
	return func(y A) B {
		return x(y)(y)
	}
}

type mc[A, B any] func(mc[A, B]) B

func M[A, B any](x mc[A, B]) B {
	return x(x)
}
