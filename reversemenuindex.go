package piscine

func ReverseMenuIndex(menu []string) []string {
	n := len(menu)
	if menu == nil || len(menu) <= 0 {
		return nil
	}
	reversedMenu := make([]string, n)
	for i := 0; i < n; i++ {
		reversedMenu[i] = menu[n-i-1]
	}
	return reversedMenu
}
