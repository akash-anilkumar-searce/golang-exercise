package main

import (
	"fmt"
)

func main() {
	defer school()
	fmt.Println("play time")
}

func school() {
	defer func() {
		fmt.Println(" school function defer ran")
	}()
	fmt.Println("school function ran")
}
