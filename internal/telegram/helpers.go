package telegram

func containString(array []string, str string) bool {
	for _, val := range array {
		if val == str {
			return true
		}
	}

	return false
}
