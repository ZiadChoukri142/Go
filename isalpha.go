package piscine

func IsAlpha(s string) bool {
	length := len(s)
	for i := 0; i < length; i++ {
		if !(s[i] >= '0' && s[i] <= '9' || s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}
	return true
}
