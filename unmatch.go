package piscine

func Unmatch(a []int) int {
	var found int
	for _, v := range a {
		found = 0
		for _, z := range a {
			if v == z {
				found++
			}
		}
		if found%2 != 0 {
			return v
		}
	}
	return -1
}
