package main

import (
	"fmt"
)

type akash int

var x akash

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}