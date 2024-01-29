package piscine

func CountIf(f func(string) bool, tab []string) int {
	numberOfElements := 0
	for _, s := range tab {
		if f(s) == true {
			numberOfElements++
		}
	}
	return numberOfElements
}
