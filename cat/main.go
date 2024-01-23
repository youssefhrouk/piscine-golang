package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := "Hello"
	if len(os.Args) < 2 {
		for _, v := range str {
			z01.PrintRune(rune(v))
		}
	}

	for _, v := range os.Args[1:] {
		content, err := ioutil.ReadFile(v)
		if err != nil {
			rr := "ERROR:"
			for _, v := range rr {
				z01.PrintRune(rune(v))
			}
			// for _, j := range content {
			// 	z01.PrintRune(rune(j))
			// }
		}

		for _, v := range content {
			z01.PrintRune(rune(v))
		}
	}
}
