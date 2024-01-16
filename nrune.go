package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	if n <= 0 || n > len(str) {
		return 0
	}
	nth := ' '
	for i := 0; i <= len(str); i++ {
		if i == n {
			nth = str[i-1]
		}
	}
	return nth
}
