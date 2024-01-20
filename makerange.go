package piscine

func MakeRange(min, max int) []int {
	length := max - min

	if max <= min {
		return nil
	}
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = i + min
	}
	return slice
}
