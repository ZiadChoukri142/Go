package piscine

func IterativeFactorial(nb int) int {
	if nb > 69 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	if nb < 0 {
		return 0
	}

	return nb * IterativeFactorial(nb-1)
}
