package piscine

func StringToIntSlice(str string) []int {
	st := []rune(str)
	nbr := []int{}
	for _, v := range st {
		nbr = append(nbr, int(v))
	}

	return nbr
}
