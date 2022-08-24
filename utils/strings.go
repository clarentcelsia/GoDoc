package utils

// Similar to when creating a new extension
type MyString string

func (m *MyString) IsExist(strs []string, str string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}
