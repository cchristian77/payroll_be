package util

// Contains will check whether Target exists in list of Source.
func Contains[T comparable](source []T, target T) bool {
	for _, item := range source {
		if item == target {
			return true
		}
	}

	return false
}
