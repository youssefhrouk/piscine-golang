package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, s1 := range a {
		for _, s2 := range s1 {
			z01.PrintRune(s2)
		}
		z01.PrintRune('\n')
	}
}
