package option

type Option[A any] struct {
	value     A
	has_value bool
}

func (o Option[A]) IsSome() bool {
	return o.has_value
}
func (o Option[A]) IsNone() bool {
	return !o.has_value
}
func (o Option[A]) UnWrap() (A, bool) {
	return o.value, o.has_value
}

func Some[A any](value A) Option[A] {
	return Option[A]{value, true}
}
func None[A any]() Option[A] {
	var zero A
	return Option[A]{zero, false}
}

func Bind[A, B any](option Option[A], f func(A) Option[B]) Option[B] {
	if option.has_value {
		return f(option.value)
	} else {
		return None[B]()
	}
}

func Map[A, B any](option Option[A], f func(A) B) Option[B] {
	if option.has_value {
		return Some(f(option.value))
	} else {
		return None[B]()
	}
}
