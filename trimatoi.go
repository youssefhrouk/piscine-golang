package piscine

func Atoi(s string) int {
	rslt := 0
	sign := 1
	for lengh, i := range s {
		if i == '+' {
			if lengh == 0 {
				sign = 1
			} else {
				return 0
			}
		} else if i == '-' {
			if lengh == 0 {
				sign = -1
			} else {
				return 0
			}
		} else if i >= '0' && i <= '9' {
			rslt = rslt*10 + int(i-'0')
		} else {
			return 0
		}
	}
	return rslt * sign
}

func TrimAtoi(s string) int {
	str := ""
	isPositive := true
	for _, value := range s {
		if str == "" && value == '-' {
			isPositive = false
		}
		if value >= '0' && value <= '9' {
			str = str + string(value)
		}
	}
	if isPositive {
		return Atoi(str)
	} else {
		return (-(Atoi(str)))
	}
}
