package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digitCounts := [10]int{}

	for n != 0 {
		digit := n % 10
		digitCounts[digit]++
		n /= 10
	}

	for i := 0; i <= 9; i++ {
		count := digitCounts[i]
		for count > 0 {
			z01.PrintRune(rune('0' + i))
			count--
		}
	}

	z01.PrintRune('\n')
}
