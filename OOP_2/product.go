package main

import(
	"fmt"
)

type product struct {
	name string
	price money
}
// now we are making methods
func (p *product)print(){
	fmt.Printf("%-15s: %s\n", p.name, p.price.string())
}

func (p *product)discount(ratio float64){
	p.price*=money((100-ratio)/100)
	fmt.Printf("The price of the %s after discount: %.2f\n", p.name, p.price)
}