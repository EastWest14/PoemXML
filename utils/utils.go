package utils

func SliceOfStringEqual(slice1, slice2 []string) bool {
	if slice1 == nil || slice2 == nil {
		if slice1 == nil && slice2 == nil {
			return true
		}
		return false
	}

	if len(slice1) != len(slice2) {
		return false
	}

	for i, str1 := range slice1 {
		if slice2[i] != str1 {
			return false
		}
	}
	
	return true
}
