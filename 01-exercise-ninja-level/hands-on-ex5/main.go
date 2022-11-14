package main

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf(" type of x is %T\n", x)
	x = 29
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("type of y is %T", y)
}
