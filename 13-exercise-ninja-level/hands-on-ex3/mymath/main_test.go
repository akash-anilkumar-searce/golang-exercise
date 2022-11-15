package mymath

import ("testing"
"fmt")

func testavg(t *testing.T){
	type avg struct{
		data[] int ,
		answer int,
	}
	tests := test{
		test([]int{10,20,30,40,},20)
		test([]int {1,2,3,4,5},2)
	}
	for _,v := range tests{
             f := CenteredAvg(v.data)
			 if f!= v.answer{
				t.Error("recieved",f,"expected",v.data)
			 }
	}
	func EgAverage(){
		f:=CenteredAvg([]int )
	}


}
func bmtesting(b *testing.B){
	for i:=0;i<b.N;i++{
		CenteredAvg([]int{1,2,3,4,10000})
	}
}
