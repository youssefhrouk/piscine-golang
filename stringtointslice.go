package piscine

func StringToIntSlice(str string) []int {
	nbr := []int{}
	for _, v := range str {
		nbr = append(nbr, int(v))
	}

	return nbr
}
