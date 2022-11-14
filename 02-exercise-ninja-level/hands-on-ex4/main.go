package main

import (
	"fmt"
)

func main() {
	m := 88
	fmt.Printf("decimal %d\t binary %b\t hexadecimal%#x\n", m, m, m)
	n := m << 1
	fmt.Printf("decimal %d\t binary %b\t hexadecimal%#x", n, n, n)
}
