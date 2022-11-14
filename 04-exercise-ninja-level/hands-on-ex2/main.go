package main

import "fmt"

func main(){
   x :=[]int {82, 83, 45, 47, 90, 87, 68, 29, 80, 71}
   for i,v := range x{
	    fmt.Println(i,v)
   }
   fmt.Printf("%T\n",x)

}