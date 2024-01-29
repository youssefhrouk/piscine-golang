package piscine

func ListForEach(f func(int), a []int) {
	for i := range a {
		f(a[i])
	}
}
