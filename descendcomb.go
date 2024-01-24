package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for v := '9'; v >= '0'; v-- {
				for k := '9'; k >= '0'; k-- {
					if (i == v && j > k) || i > v {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(v)
						z01.PrintRune(k)
						if i != '0' || j != '1' || v != '0' || k != '0' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
