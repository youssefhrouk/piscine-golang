package piscine

func	UltimateDivMod(a *int, b *int) {
    nb := *a
	*a = *a / *b
	*b = nb % *b
}
/*func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}*/
