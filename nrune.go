package piscine

func NRune(s string, n int) rune {
	if n > len(s) || n <= 0 {
		return 0
	}
	runes := []rune(s)
	return runes[n-1]
}
