package piscine

func Unmatch(a []int) int {
	length := len(a)

	for i := 0; i < length; i++ {
		count := 0
		for j := 0; j < length; j++ {
			if a[i] == a[j] {
				count++
			}
		}

		if count%2 != 0 {
			return a[i]
		}
	}

	return -1
}
