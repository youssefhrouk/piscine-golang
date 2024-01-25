package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)

	for i := 0; i < length/2; i++ {
		menu[i], menu[length-i-1] = menu[length-i-1], menu[i]
	}

	return menu
}
