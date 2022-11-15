//testing the go program

package dog

import ("testing"
 "fmt")


func TestYears(t *testing.T){
    n:=Years(10)
	if n!=70{
		t.Error("recieved",n,"expected",70)
	}
}
func TestYearsTwo(t *testing.T){
	n:=YearsTwo(10)
	if n!=70{
		t.Error("recieved",n,"expected",70)
	}
}

func EgYears(){
	fmt.Println(Years(20))
	//OUTPUT:
	//140

}
func EgYearsTwo(){
	fmt.Println(YearsTwo(20))
	//OUTPUT:
	//140
}
func BMYears(b *testing.B){
	for i:=0;i<b.N;i++{
		Years(50)
	}
}

func BMYearstwo(b *testing.B){
	for i:=0;i<b.N;i++{
		YearsTwo(50)
	}
}

