package piscine

func DescendAppendRange(max, min int) []int {
	name := []int{}
	for x := max; x > min; x-- {
		name = append(name, x)
	}
	return name
}
