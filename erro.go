package erro

func Recover(fn func(error)) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			fn(err)
		} else {
			panic(r)
		}
	}
}

func Check0(err error) {
	if err != nil {
		panic(err)
	}
}

func Check1[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Check2[S, T any](s S, t T, err error) (S, T) {
	if err != nil {
		panic(err)
	}
	return s, t
}

func Check3[R, S, T any](r R, s S, t T, err error) (R, S, T) {
	if err != nil {
		panic(err)
	}
	return r, s, t
}
