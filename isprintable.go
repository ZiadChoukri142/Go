package piscine

func IsPrintable(s string) bool {
	length := len(s)
	for i := 0; i < length; i++ {
		if !(s[i] >= 32 && s[i] <= 127) {
			return false
		}
	}
	return true
}
