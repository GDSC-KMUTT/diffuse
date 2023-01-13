package value

import "backend/types/response"

func Contain[T comparable](arr []T, elem T) bool {
	for _, a := range arr {
		if a == elem {
			return true
		}
	}
	return false
}

func ContainVal[T comparable](arr []*T, elem *T) bool {
	for _, a := range arr {
		if *a == *elem {
			return true
		}
	}
	return false
}

func Map[A any, B any](a []A, mapper func(a A) (B, *response.ErrorInstance)) ([]B, *response.ErrorInstance) {
	result := make([]B, 0)
	for _, el := range a {
		mapped, err := mapper(el)
		if err != nil {
			return nil, err
		}
		result = append(result, mapped)
	}
	return result, nil
}
