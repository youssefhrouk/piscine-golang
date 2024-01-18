package main

import (
	"fmt"
)

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
		if value == '-' && str == "" {
			isPositive = false
		}
		if value >= '0' && value <= '9' {
			str = str + string(value)
		}
	}
	if isPositive {
		return Atoi(str)
	} else {
		return -Atoi(str)
	}
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
