package main
import "fmt"


func main(){
	for i:=10;i<=100;i++{
		fmt.Println("the remainder of %v when divided by 4 is %v",i,i%4)
	}
}