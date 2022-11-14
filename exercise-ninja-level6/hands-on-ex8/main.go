package main

import "fmt"

func main() {
	f := prog()
	fmt.Println(f())
}

func prog() func() int {
	return func() int {
		return 33
	}
}
