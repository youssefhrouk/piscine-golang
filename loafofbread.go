package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var res string
	i := 0

	for i < len(str) {
		if str[i] == ' ' {
			i++
			continue
		}
		res += string(str[i])
		i += 2

		if len(res) == 5 {
			res += "\n"
			res = ""
		}
	}

	return res
}
