package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Printf("File name missing")
	} else if len(arg) > 1 {
		fmt.Printf("Too many arguments")
	} else {
		file, err := ioutil.ReadFile("quest8.txt")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(file))
	}
}
