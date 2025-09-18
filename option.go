package option

import result "github.com/SamsPrograms/result"

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

func ToResult[A any](option Option[A], err error) result.Result[A] {
	if option.has_value {
		return result.Ok(option.value)
	} else {
		return result.Err[A](err)
	}
}

func FromResult[A any](result result.Result[A]) Option[A] {
	if value, err := result.Unwrap(); err != nil {
		return Some(value)
	} else {
		return None[A]()
	}
}
