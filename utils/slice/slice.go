package slice

func Contain[T comparable](s []T, p T) bool {
	for _, ss := range s {
		if p == ss {
			return true
		}
	}
	return false
}

func Unique[T comparable](s []T) (res []T) {
	seen := make(map[T]bool)
	for _, ss := range s {
		if !seen[ss] {
			res = append(res, ss)
			seen[ss] = true
		}
	}
	return
}