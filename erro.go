package erro

func Recover(fn func(any)) {
	if r := recover(); r != nil {
		fn(r)
	}
}

func Check[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
