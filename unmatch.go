package piscine

func Unmatch(a []int) int {
	index := 0

	for i, v := range a {
		ismatch := false
		for j, w := range a {
			if i != j && v == w {
				ismatch = true
				break
			}
		}
		if !ismatch {
			index = v
			break
		}
	}
	return index
}
