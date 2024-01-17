package piscine

func ToUpper(s string) string {
	str := ""
	for _, v := range s {
		if v >= 97 && v <= 122 {
			str += string(v - 32)
		} else {
			str += string(v)
		}
	}
	return str
}
