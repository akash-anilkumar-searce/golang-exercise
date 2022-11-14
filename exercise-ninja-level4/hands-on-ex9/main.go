package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`akash`:      []string{`ice`, `coffee`, `loyal`},
		`sabhari`: []string{`india`, `sports`, `Cinema`},
		`dr`:         []string{`sincere`, `cream`, `Sunsets`},
	}

    m[`subramanyam`]=[]string{`happy`,`peaceful`,`fun`}
  
	for a, v := range m {
		fmt.Println("This is the record for", a)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}