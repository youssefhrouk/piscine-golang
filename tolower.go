package piscine

func ToLower(s string) string {
	str := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			str += string(v + 32)
		} else {
			str += string(v)
		}
	}
	return str
}
