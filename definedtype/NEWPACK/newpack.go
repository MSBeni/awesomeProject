package newpack

import "fmt"


// this is a test for the type and undelting type
func Newpack(){
	type duration float32
	var newvalue duration = 66.31
	fmt.Println("That's a new package", newvalue)
}