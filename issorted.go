package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := range a {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}
