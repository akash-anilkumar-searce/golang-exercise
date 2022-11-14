package main

import (
	"fmt"
)

func main() {
	m := 88
	fmt.Printf("%d\t%b\t%#x\n", m, m, m)
	n := m << 1
	fmt.Printf("%d\t%b\t%#x", n, n, n)
}
