package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i := range strs {
		str += strs[i]
		if i < len(strs)-1 {
			str += sep
		}
	}
	return str
}
