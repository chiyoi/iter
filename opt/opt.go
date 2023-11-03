package opt

type Option[A any] struct {
	v    A
	some bool
}

func Some[A any](v A) Option[A] {
	return Option[A]{v, true}
}

func None[A any]() Option[A] {
	return Option[A]{}
}

func Then[A, B any](a Option[A], f func(A) Option[B]) Option[B] {
	if !a.some {
		var zero B
		return Option[B]{zero, false}
	}
	return f(a.v)
}

func Or[A any](a Option[A], f func() Option[A]) Option[A] {
	if a.some {
		return a
	}
	return f()
}
