package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Hello")
	}

	for _, v := range os.Args[1:] {
		content, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Printf("ERROR%s: %s\n", content, err)
			return
		}

		fmt.Println(string(content))
	}
}
