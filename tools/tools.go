package tools

func SliceContainString(s []string, p string) bool {
	for _, ss := range s {
		if ss == p {
			return true
		}
	}
	return false
}
