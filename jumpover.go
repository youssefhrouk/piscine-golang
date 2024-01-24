package piscine

func JumpOver(str string) string {
	st := ""

	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		st += string(str[i])
	}
	st += "\n"
	return st
}
