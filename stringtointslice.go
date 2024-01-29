package piscine

func StringToIntSlice(str string) []int {
	var name []int
	for _, x := range str {
		name = append(name, int(x))
	}
	return name
}
