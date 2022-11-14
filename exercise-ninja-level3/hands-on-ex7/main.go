package main

import (
	"fmt"
)

func main() {
	a := "Akash"

	if a == "Akash" {
		fmt.Println(a)
	} else if a == "Sabhari" {
		fmt.Println("friend1", a)
	} else {
		fmt.Println("neither")
	}
}
