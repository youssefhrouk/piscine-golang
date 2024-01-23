package piscine

func Any(f func(string) bool, a []string) bool {
	ip := true
	for _, v := range a {
		ip = f(string(v))
	}
	return ip
}
