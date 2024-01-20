package piscine

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	index := 0

	for i := 0; i <= len(s); i++ {
		if i == len(s) || isWhitespace(s[i]) {
			if i > index {
				result = append(result, s[index:i])
			}
			index = i + 1
		}
	}
	return result
}
