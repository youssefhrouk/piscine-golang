package piscine

func BasicAtoi(s string) int {
	cp := 0
	for _, i := range s {
		cp = cp*10 + int(i) - '0'
	}
	return cp
}

/*func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}*/
