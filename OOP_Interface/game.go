package main

import(
	"fmt"
)

type game struct{
	name string
	price money
	countries map[string]bool
}

func (g game) print(){
	fmt.Printf("the name od the game is %s and the price is %s which is allowed in these countries: %v \n", g.name, g.price.print(), g.countries )
}

func (g *game) discount(ratio money){
	g.price*=1-ratio
}