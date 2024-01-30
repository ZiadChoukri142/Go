package piscine

func BubbleSort2(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func Max(a []int) int {
	sortedArray := BubbleSort2(a)
	length := len(sortedArray)
	return sortedArray[length-1]
}
