package values

func IsLessThanZero(value *int) bool {
	if value == nil {
		return false
	}

	return *value < 0
}
