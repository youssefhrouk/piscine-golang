package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ssa := os.Args[1:]
	tr := len(ssa)
	for i := 0; i < tr-1; i++ {
		for j := i + 1; j < tr; j++ {
			if ssa[i] > ssa[j] {
				ssa[i], ssa[j] = ssa[j], ssa[i]
			}
		}
	}
	for _, arg := range ssa {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
