package main

import "fmt"

func main(){

	c:=make(chan int)
	go func (){
		c<-98
	}()
	fmt.Println(<-c)

	
}