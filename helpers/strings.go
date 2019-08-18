package helpers

// SliceContainsString returns idx and true if a found in s. Otherwise -1 and false.
func SliceContainsString(s []string, a string) bool {
	_, r := SliceContainsStringAndPos(s, a)

	return r
}

// SliceContainsStringAndPos returns idx and true if a found in s. Otherwise -1 and false.
func SliceContainsStringAndPos(s []string, a string) (int, bool) {
	for i, b := range s {
		if b == a {
			return i, true
		}
	}

	return -1, false
}
