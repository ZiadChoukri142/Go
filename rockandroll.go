package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if n == 1 {
		return "error: non divisible"
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return "error: non divisible"
		}
	}

	switch {
	case n%2 == 0 && n%3 == 0:
		return "rock and roll \n"
	case n%2 == 0:
		return "rock \n"
	case n%3 == 0:
		return "roll \n"
	default:
		return " "
	}
}
