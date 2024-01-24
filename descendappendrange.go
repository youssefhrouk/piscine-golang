package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		return []int{}
	}
	var r []int
	for i := max; i > min; i-- {
		r = append(r, i)
	}
	return r
}
