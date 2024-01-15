package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	m := 1
	for i := 1; i <= power; i++ {
		m = m * nb
	}
	return m
}

/* func main() {
 	fmt.Println(IterativePower(4, 2))
 }
