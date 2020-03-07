package main

//import(
//	"fmt"
//)

type book struct {
	product
	publish interface{}
}

func (b *book)print(){
	b.product.print()
	//p:= format(b.publish)
	//fmt.Printf("%-15s: %s\n", p)
}

