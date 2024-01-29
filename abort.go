package piscine

func BubbleSorter(arr []int) {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	BubbleSorter(arr)

	result := arr[len(arr)/2]
	return result
}
