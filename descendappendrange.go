package piscine

func DescendAppendRange(max, min int) []int {
	var slice []int
	for i := max; i > min; i-- {
		if max > min {
			slice = append(slice, i)
		}
	}
	return slice
}
