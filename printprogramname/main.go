package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main(){
	for i,v := range os.Args[0]{
		if i>1{
		z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}