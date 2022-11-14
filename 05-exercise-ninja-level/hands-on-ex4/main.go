package main

import (
	"fmt"
)

func main() {

	a := struct {
		name    string
		bestfriend  map[string]int
		food []string
	}{
		name: "akash",
		bestfriend: map[string]int{
			"sabhari": 18,
			"subramanyam":      77,
			"aswin":          809,
		},
		food: []string{
			"dosa",
			"pan",
		},
	}
	fmt.Println(a.name)
	fmt.Println(a.bestfriend)
	fmt.Println(a.food)

	for k, v := range a.bestfriend {
		fmt.Println(k, v)
	}

	for i, val := range a.food {
		fmt.Println(i, val)
	}
}
