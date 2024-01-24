package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	numPlayers := len(deck) / 3 // Number of players based on deck size

	for i := 0; i < numPlayers; i++ {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		z01.PrintRune(rune(49 + i)) // Printz01. player number
		z01.PrintRune(':')
		z01.PrintRune(' ')

		// Print player's cards
		for j := 0; j < 3; j++ {
			z01.PrintRune(rune(deck[i*3+j] + 48)) // Print card value as ASCII digit
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune('\n')
	}
}
