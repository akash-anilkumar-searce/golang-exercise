package main

import (
	"fmt"
)

const (
	w = 2017 + iota
	x
	y
	z
)

func main() {
	fmt.Println("value after using iota ",w)
	fmt.Println("value after using iota ",x)
	fmt.Println("value after using iota ",y)
	fmt.Println("value after using iota ",z)
}
