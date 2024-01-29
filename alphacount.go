package piscine

func AlphaCount(s string) int {
	length := len(s)
	finalLength := 0
	for i := 0; i < length; i++ {
		if (s[i] >= 65) && (s[i] <= 90) || (s[i] >= 97) && (s[i] <= 122) {
			finalLength++
		}
	}
	return finalLength
}
