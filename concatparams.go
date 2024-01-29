package piscine

func ConcatParams(args []string) string {
	var name string
	x := 0

	for range args {
		x++
	}
	for i, v := range args {
		name += v
		if i != x-1 {
			name += "\n"
		}
	}
	return name
}
