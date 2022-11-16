package word

import (
	"strings"

	"13-exercise-hands-on-ex2/quote"
)

func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

func Count(s string) int {
	m := UseCount(quote.SunAlso)

	return m[s]
}