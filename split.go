package piscine

func Split(s, sep string) []string {
	var result []string
	index := 0

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i:i+len(sep)] == sep {
			if i > index {
				result = append(result, s[index:i])
			}
			index = i + len(sep)
		}
	}
	return result
}