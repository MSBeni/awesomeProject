package main

import(
	"fmt"
	"strconv"
	"awesomeProject/arraytest/arrayofarray"
)

func main(){
	var books[4] string

	for i,_:= range books{
		books[i] = "Book" + strconv.Itoa(i+1)
	}

	for _,v := range books{
		fmt.Printf("The books names: %#v\n", v)
	}
	for _, v := range [...]string{"A", "v", "w"}{
		fmt.Printf("The letters names: %#v\n", v)
	}
	arrayofarray.MedianOG()
}