package piscine

func StringToIntSlice(str string) []int {
	nbr := []int{}
	if str == "" {
		return nil
	}
	for _, v := range str {
		nbr = append(nbr, int(v))
	}

	return nbr
}
