package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := 1; i < len(arg); i++ {
		slice := []rune(arg[i])
		for _, j := range slice {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')

	}
}
