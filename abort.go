package piscine

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				tmp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = tmp

			}
		}
	}
	return slice[2]
}
