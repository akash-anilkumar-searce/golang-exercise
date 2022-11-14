package main

import "fmt"

type Person struct {
	name string
}

type Human interface {
	speak()
}

func (p *Person) speak() {
	fmt.Println(p.name + " is speaking!")
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	Appu := Person{"Akash"}

	saySomething(&Appu)

	
}