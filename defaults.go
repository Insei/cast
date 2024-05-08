package cast

func Default[T supported]() T {
	return *new(T)
}
