package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	numP := 4
	cp := len(deck) / numP

	for i := 0; i < numP; i++ {

		playerC := deck[i*cp : (i+1)*cp]

		fmt.Printf("Player %d: %v$\n", i+1, playerC)
	}
}
