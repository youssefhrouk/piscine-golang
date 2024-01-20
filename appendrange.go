package piscine

func AppendRange(min, max int) []int {
	var slice []int

	for i := min; i < max; i++ {
		if max > min {
			slice = append(slice, i)
		}
	}
	return slice
}
