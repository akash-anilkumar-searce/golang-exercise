package main

import (
	"fmt"
)

func main() {

	func() {
		for i := 10; i < 90; i++ {
			fmt.Println("anonymous function",i)
		}
	}()

	fmt.Println("done")
}
