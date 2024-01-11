package piscine

func BasicAtoi2(s string) int {
	a := []rune(s)
	b := len(s)
	c := 0
	for i := 0; i < b; i++ {
		if a[i] < '0' || a[i] > '9' {
			return 0
		} else {
			c *= 10
			c += int(a[i]) - '0'
		}
	}
	return c
}
