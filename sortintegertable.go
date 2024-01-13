package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := len(table) - 1; j > i; j-- {
			if table[j] > table[i] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
