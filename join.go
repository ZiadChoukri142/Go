package piscine

func Join(strs []string, sep string) string {
	name := ""
	for i, res := range strs {
		name += res
		if i != len(strs)-1 {
			name += sep
		}
	}
	return name
}
