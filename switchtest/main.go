package main

import (
	"awesomeProject/switchtest/forlooptest"
	"fmt"
	"strings"
	"time"
)

func main(){

	//fmt.Printf("%T \n %d\n",a, a)
	words := strings.Fields("You Such A Crazy Boy")
	switch a:= time.Now().Hour(); {
	case a>=18:
		fmt.Println("Good Evening")
	case a>=12:
		fmt.Println("Good Afternoon")
	case a>=6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Night")
	}
	forlooptest.Forlooptest()
	for i,v:= range words{
		fmt.Printf("#%-2d: %q\n", i+1,  v)
	}

}