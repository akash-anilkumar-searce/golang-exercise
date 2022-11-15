//testing the go program

package dog

import "testing"

type iteration struct{
	n int
}

func BMYears(t *testing.T){
	for i:=0;i<t.n;i++{
		Years(50)
	}
}

func BMYearstwo(t *testing.T){
	for i:=0;i<t.n;i++{
		YearsTwo(22)
	}
}

