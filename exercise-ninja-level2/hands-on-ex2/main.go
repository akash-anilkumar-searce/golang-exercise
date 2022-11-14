package main

import "fmt"


func main() {
	g := (59 == 59)
	h := (59 <= 68)
	i := (59 >= 68)
	j := (59 != 68)
	k := (59 < 68)
	l := (59 > 68)
	fmt.Println(g, h, i, j, k, l)
}