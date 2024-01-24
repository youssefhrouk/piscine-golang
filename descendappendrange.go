package main

import (
	"fmt"
)

func DescendAppendRange(max, min int) []int {
	var slice []int
	if max <= min {
		return nil
	}
	for i := max; i > min; i-- {
		slice = append(slice, i)
	}
	return slice
}

func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
}
