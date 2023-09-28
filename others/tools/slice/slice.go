package slice

func ContainString(s []string, p string) bool {
	for _, ss := range s {
		if ss == p {
			return true
		}
	}
	return false
}
