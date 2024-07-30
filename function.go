package goutil

import "reflect"

func Ptr[T any](v T) *T {
	return &v
}

func Unptr[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func First[T any](elems ...T) T {
	if len(elems) == 0 {
		panic("First: empty slice")
	}
	for i := 0; i < len(elems); i++ {
		if !reflect.ValueOf(elems[i]).IsZero() {
			return elems[i]
		}
	}
	return elems[0]
}

func Last[T any](elems ...T) T {
	if len(elems) == 0 {
		panic("Last: empty slice")
	}
	for i := len(elems) - 1; i >= 0; i-- {
		if !reflect.ValueOf(elems[i]).IsZero() {
			return elems[i]
		}
	}
	return elems[0]
}

func If[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}

func Map[T, U any](elems []T, fn func(T) U) []U {
	ret := make([]U, len(elems))
	for i := 0; i < len(elems); i++ {
		ret[i] = fn(elems[i])
	}
	return ret
}

func Reduce[T, U any](elems []T, init U, fn func(U, T) U) U {
	ret := init
	for i := 0; i < len(elems); i++ {
		ret = fn(ret, elems[i])
	}
	return ret
}

func Filter[T any](elems []T, fn func(T) bool) []T {
	ret := make([]T, 0)
	for i := 0; i < len(elems); i++ {
		if fn(elems[i]) {
			ret = append(ret, elems[i])
		}
	}
	return ret
}

func Unique[T comparable](elems []T) []T {
	ret := make([]T, 0, len(elems))
	t := make(map[T]struct{}, len(elems))
	for i := 0; i < len(elems); i++ {
		if _, ok := t[elems[i]]; !ok {
			t[elems[i]] = struct{}{}
			ret = append(ret, elems[i])
		}
	}
	return ret
}
