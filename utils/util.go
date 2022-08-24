package utils

// Utility for search specific item in list
func Contains(list []interface{}, data interface{}) bool {
	for _, v := range list {
		if v == data {
			return true
		}
	}
	return false
}
