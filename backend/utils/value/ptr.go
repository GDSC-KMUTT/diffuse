package value

func Ptr[T any](v T) *T {
	return &v
}

func Val[T any](pointer *T, defaultValue ...T) T {
	var variable T
	if pointer != nil {
		variable = *pointer
	} else if len(defaultValue) > 0 {
		variable = defaultValue[0]
	}
	return variable
}

func Empty[T comparable](pointer *T) bool {
	var def T
	if pointer == nil || *pointer == def {
		return true
	}
	return false
}
