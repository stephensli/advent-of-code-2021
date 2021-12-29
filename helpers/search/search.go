package search

func StringArrayContains(input []string, value string) bool {
	for _, s := range input {
		if s == value {
			return true
		}
	}

	return false
}
