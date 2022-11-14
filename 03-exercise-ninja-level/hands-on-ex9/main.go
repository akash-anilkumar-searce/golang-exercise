package main

import (
	"fmt"
)

func main() {
	favsport :="basketball"
	switch favsport{
	case "football":
		fmt.Println("got to stadium")
	case "basketball":
		fmt.Println("go to basketball court")
	case "swimmming":
		fmt.Println("go to pool")
	}
}