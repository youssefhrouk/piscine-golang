package piscine

func DescendAppendRange(max, min int) []int {
	var slice []int
	if max <= min {
		return nil
	}
	for i := max; i > min; i-- {
		slice = append(slice, i)
	}
	return slice
}