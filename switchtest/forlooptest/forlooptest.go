package forlooptest

import(
	"fmt"
	"os"
)
// just a simple code
func Forlooptest(){
	a := os.Args
	for i:=1; i<len(a); i++{
		//fmt.Printf("%s\n", a[i])
		for j:=0; j<len(a[i]);j++{
			fmt.Println(string(a[i][j]))
		}
	}
}