package main

import "fmt"

type Book struct{
	name string
	price money
}

func (b Book) print(){
	fmt.Printf("the name of the BOOK is %s and the Price is %s \n", b.name, b.price.price_game())
}