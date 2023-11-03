package opt

type Option[A any] struct {
	v    A
	some bool
}

type NoneType struct{}

func Some[A any](v A) Option[A] {
	return Option[A]{v, true}
}

func None() Option[NoneType] {
	return Option[NoneType]{struct{}{}, false}
}

func Then[A, B any](a Option[A], f func(A) Option[B]) Option[B] {
	if !a.some {
		var zero B
		return Option[B]{zero, false}
	}
	return f(a.v)
}
