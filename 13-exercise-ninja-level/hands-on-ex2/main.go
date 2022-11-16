package main

import (
	"fmt"
  "13-exercise-hands-on-ex2/word"
  "13-exercise-hands-on-ex2/quote"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}