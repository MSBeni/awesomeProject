package shortifpackage

import (
	"fmt"
	"os"
	"strconv"
)
// receive a number and double it
func Shortifpack(){
	var n int
	if a := os.Args; len(a)<3 {
		fmt.Println("Please enter a number besides the username and password")
	} else if n, err := strconv.Atoi(a[3]); err!=nil{
		fmt.Printf("Cannot Convert %q. \n", a[3])
	}else{
		fmt.Printf("%s *2 %d\n", a[3], n*2)
	}
	fmt.Println(n)
}
