package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	compactSlice := make([]string, 0)
	for _, value := range slice {
		if value != "" {
			compactSlice = append(compactSlice, value)
		}
	}
	*ptr = compactSlice
	return len(compactSlice)
}
