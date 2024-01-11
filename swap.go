package piscine

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

/*func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}*/
