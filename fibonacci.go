package piscine

func Fibonacci(index int) int {
	if index < 2 {
		return index
	}
	return (Fibonacci(index-1) + Fibonacci(index-2))
}

// func main() {
// 	arg1 := 4
// 	fmt.Println(Fibonacci(arg1))
// }
