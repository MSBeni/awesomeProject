package main

import "fmt"

type printer interface {
	print()
}

type List []printer

func (l List) print(){
	if len(l)==0{
		fmt.Println("Sorry there is nothing in your list to be printed")
		return
	}
	for _, el := range l{
		el.print()
	}
}
