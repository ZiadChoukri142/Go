package piscine

func ToUpper(s string) string {
	result := []rune(s)
	for i := 0; i < len(result); i++ {
		if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = result[i] - 'a' + 'A'
		}
	}
	return string(result)
}
