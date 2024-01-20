package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	var Currentmot string

	for i := 0; i < len(s); i++ {
		str := s[i]

		if str == ' ' || str == '\t' || str == '\n' || str == '\r' {
			if Currentmot != "" {
				res = append(res, Currentmot)
				Currentmot = ""
			}
		} else {
			Currentmot += string(str)
		}
	}

	if Currentmot != "" {
		res = append(res, Currentmot)
	}

	return res
}
