package piscine

func CountIf(f func(string) bool, tab []string) int {
	t := 0
	for _, v := range tab {
		if f(v) {
			t++
		}
	}
	return t
}
