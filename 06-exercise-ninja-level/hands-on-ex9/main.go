package main

import (
	"fmt"
)

func main() {

	a := func(ai []int) int {
		if len(ai) == 0 {
			return 0
		}
		if len(ai) == 1 {
			return ai[0]
		}
		return ai[0] + ai[len(ai)-1]
	}

	b := foo(a, []int{16, 23, 53, 24, 85, 65})
	fmt.Println(b)
}

func foo(f func(ai []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}