package piscine

func UltimateDivMod(a *int, b *int) {
	f := *a / *b
	e := *a % *b
	*a = f
	*b = e
}
