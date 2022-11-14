package main

import (
	"fmt"
)

func main() {
	xx := []int{891, 2234, 243, 4, 5532,326, 3527, 8,5329}
	n := foo(xx...)
	fmt.Println(n)
	
	xx2 := []int{891, 2234, 243, 4, 5532,326, 3527, 8,5329}
	n2 := bar(xx2)
	fmt.Println(n2)
}

func foo(ai ...int) int {
	total := 0
	for _, v := range ai {
		total += v
	}
	return total
}

func bar(ai []int) int {
	total := 0
	for _, v := range ai {
		total += v
	}
	return total
}