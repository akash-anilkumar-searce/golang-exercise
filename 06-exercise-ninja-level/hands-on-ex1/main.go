package main

import (
	"fmt"
)

func main() {
	a := foo()
	x, y := bar()

	fmt.Println(a, x, y)
}

func foo() int {
	return 33
}

func bar() (int, string) {
	return 2000, "great gatsby"
}

