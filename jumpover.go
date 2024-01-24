package main

import "fmt"

func JumpOver(str string) string {
	st := " "

	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		st += string(str[i])
	}
	st += "\n"
	return st
}

func main() {
	fmt.Print(JumpOver("1010101010"))
	fmt.Print(JumpOver(""))
	fmt.Print(JumpOver("t w e l v e"))
	fmt.Print(JumpOver("12"))
}
