package piscine

import "github.com/01-edu/z01"

func BubbleSort(r []rune) {
	n := len(r)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	resultRunes := []rune{}
	for n > 0 {
		digit := n % 10
		n /= 10
		resultRunes = append([]rune{rune('0' + digit)}, resultRunes...)
	}
	BubbleSort(resultRunes)
	for _, r := range resultRunes {
		z01.PrintRune(r)
	}
}
