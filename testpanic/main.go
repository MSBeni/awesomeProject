package main

import(
	"fmt"
)

func main(){
	fmt.Println("Hello")
	testpanic()
	fmt.Println("World")
}

func testpanic(){
	defer func() {
		if err := recover(); err!=nil{
			fmt.Println(err)
			fmt.Println("We recover from a panic!!!")
		}
	}()

	panic("There was an error there buddy")
}