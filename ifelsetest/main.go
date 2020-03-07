package main

import (
	"fmt"
	"os"
	//"strconv"
	"awesomeProject/ifelsetest/shortifpackage"
)

const usn = "Benji"
const pass = "1234"

func main(){
	if len(os.Args)<3 {
		fmt.Println("Please enter Your username and password to enter: ")
		return
	}
	arg := os.Args[1]
	arg2 := os.Args[2]
	if (arg == usn) && (arg2 == pass){
		fmt.Println("You enter Correctly")
	} else if (arg == usn) && (arg2 != pass){
		fmt.Println("Your Pass is not Correct!!! ")

	//if s, err := strconv.ParseFloat(arg, 64); err == nil {
	//	fmt.Printf("%T, %v\n", s, s*0.63)
	} else if (arg != usn) && (arg2 != pass){
		fmt.Println("Your Username is not Correct!!! ")
	} else{
		fmt.Println("Wrong Username or Pass!!! ")
	}
	shortifpackage.Shortifpack()
}
