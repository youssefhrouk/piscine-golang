package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	isprime := true
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			isprime = false
		}
	}
	return isprime
}
