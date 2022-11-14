package main

import (
	"fmt"
)

func main() {
	as1 := []string{"James", "Bond", "Shaken, not stirred"}
	as2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(as1)
	fmt.Println(as2)

	aas := [][]string{as1, as2}
	fmt.Println(aas)

	for i, as := range aas {
		fmt.Println("record: ", i)
		for j, val := range as {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
