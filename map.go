package piscine

func Map(f func(int) bool, a []int) []bool {
	slice := make([]bool, len(a))
	for i, v := range a {
		slice[i] = f(v)
	}
	return slice
}
