package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nb int) {
	c := '0'
	if nb == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= nb%10; i++ {
		c++
	}
	if nb/10 != 0 {
		PrintNbr(nb / 10)
	}
	z01.PrintRune(c)
}

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	sr := "x = "
	for _, i := range sr {
		z01.PrintRune(i)
	}
	PrintNbr(points.x)
	srt := ", y = "
	for _, j := range srt {
		z01.PrintRune(j)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
