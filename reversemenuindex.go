package piscine

func ReverseMenuIndex(menu []string) []string {
	name := make([]string, len(menu))
	for x := 0; x < len(menu); x++ {
		name[len(menu)-1-x] = menu[x]
	}
	return name
}
