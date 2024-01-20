package piscine

func ConcatParams(args []string) string {
	str := []string(args)
	tab := ""
	for i := range str {

		tab += str[i]
		if i < len(str)-1 {
			tab += "\n"
		}
	}
	return tab
}
