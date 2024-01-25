package piscine

func ShoppingListSort(slice []string) []string {
	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 {
			if len(slice[j-1]) > len(slice[j]) {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j--
		}
	}
	return slice
}
