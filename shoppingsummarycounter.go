package piscine

func toLower(char byte) byte {
	if 'A' <= char && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func ShoppingSummaryCounter(str string) map[string]int {
	var wordStart int
	summary := make(map[string]int)

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '.' || str[i] == ',' {
			if wordStart < i {
				word := str[wordStart:i]
				wordLower := ""
				for j := 0; j < len(word); j++ {
					wordLower += string(toLower(word[j]))
				}
				summary[wordLower]++
			}
			wordStart = i + 1
		}
	}

	if wordStart < len(str) {
		word := str[wordStart:]
		wordLower := ""
		for j := 0; j < len(word); j++ {
			wordLower += string(toLower(word[j]))
		}
		summary[wordLower]++
	}

	return summary
}
