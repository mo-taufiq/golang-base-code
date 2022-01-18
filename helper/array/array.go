package array

func RemoveDuplicateArrayOfString(arr []string) []string {
	result := []string{}

	for _, a := range arr {
		isExist := IsArrayOfStringInclude(result, a)
		if !isExist {
			result = append(result, a)
		}
	}
	return result
}

func IsArrayOfStringInclude(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}
