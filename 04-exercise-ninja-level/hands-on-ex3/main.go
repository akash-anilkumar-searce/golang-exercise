package main

import (
	"fmt"
)

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("selected element of slice are",a[:5])
	fmt.Println("selected element of slice are",a[5:])
	fmt.Println("selected element of slice are",a[2:7])
	fmt.Println("selected element of slice are",a[1:6])
	fmt.Println("print full slice",a)
}

