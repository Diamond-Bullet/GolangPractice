package slice

func Contain[T comparable](s []T, p T) bool {
	for _, ss := range s {
		if p == ss {
			return true
		}
	}
	return false
}
