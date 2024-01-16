package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	for i := 0; i <= len(a)-len(b); i++ {
		if toFind == s[i:i+len(b)] {
			return i
		}
	}
	return -1
}
