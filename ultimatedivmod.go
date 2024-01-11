package piscine

func UltimateDivMod(a *int, b *int) {
	nb := *a
	*a = *a / *b
	*b = nb % *b
}
