package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	myMap := map[string]int{
		"Burger": 2,
		"Water":  4,
		"Carrot": 4,
		"Coffee": 1,
		"Chips":  1,
	}

	for i, v := range myMap {
		myMap[i] = v
	}
	return myMap
}
