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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("######")
	}

}