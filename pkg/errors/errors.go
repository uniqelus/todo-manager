package liberrors

func Try(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	Try(err)
	return v
}
