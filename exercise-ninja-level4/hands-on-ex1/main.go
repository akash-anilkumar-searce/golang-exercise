package main

import (
	"fmt"
)

func main() {
	a := [5]int{88, 3, 54, 84, 26}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}
