package mymath

import ("testing"
"fmt")

func TestCenteredAvg(t *testing.T){
	type test struct{
		data []int
		answer float64
	}
	tests := []test{test{[]int{10,20,30,40},25},
	    test{[]int{1,2,3,4,},2.5,}}
	for _,v :=range tests{
		f := CenteredAvg(v.data)
		if f != v.answer{
                 t.Error("recieved",f,"expected",v.answer)
		}
	}
}

func ExampleCenteredAvg(){
	fmt.Println(CenteredAvg([]int{10,20,30,40}),20)
	//OUTPUT 
	//20
}

func BenchmarkCenteredAvg(b *testing.B){
	for i:=0;i<b.N;i++{
		CenteredAvg([]int{100,200,300,400,500})
	}
}



func bmtesting(b *testing.B){
	for i:=0;i<b.N;i++{
		CenteredAvg([]int{1,2,3,4,10000})
	}
}
