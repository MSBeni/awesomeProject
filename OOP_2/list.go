package main

import "fmt"


type printer interface {
	print()
}
type list []printer

func (l list) print(){
	if len(l)==0{
		fmt.Println("Sorry we are waiting for a value to be entered ...")
	}
	for _,v:= range l{
		v.print()
	}
}

func (l list) discount(ratio float64){
	type discounter interface {
		discount(float64)
	}
	for _, val:= range l{
		// g, ok := val.interface{discount(float64)}
		if g, ok:= val.(discounter); ok{
			g.discount(ratio)
		}
		//g, ok := val.(discounter)
		//if !ok{
		//	continue
		//}
		//g.discount(ratio)
	}
}