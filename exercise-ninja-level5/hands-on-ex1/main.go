package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "Akash",
		last:  "Anilkumar",
		favFlavors: []string{
			"butterscotch",
			"pistah",
			"mango",
		},
	}

	p2 := person{
		first: "sabhari",
		last:  "r",
		favFlavors: []string{
			"chocolate",
			"dark chocolate",
			"capuccino",
		},
	}

	fmt.Println("first name"p1.first)
	fmt.Println("last name"p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println("first name"p2.first)
	fmt.Println("last name"p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}