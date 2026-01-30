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

func Check[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
