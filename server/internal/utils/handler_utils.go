package utils

func ErrorAsMap(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}
