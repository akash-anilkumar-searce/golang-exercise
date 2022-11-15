package main

import ("fmt"
"hands-on-ex1/dog")


type canine struct{
	name string
	age int
}

func main(){
	p1:= canine{name:"hi",age:dog.Years(10),}
	fmt.Println(p1)
}