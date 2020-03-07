package main

import "fmt"

func main(){

	store:= list{
		&book{product{name:  "History of Iran", price: 1000}, nil},
		&game{product{name:  "Hitman", price: 100,}},
		&game{product{name:  "Tetris", price: 50,}},
		&puzzle{product{name:"purepuz", price: 30,}},
		&toy{product{name: "Mench", price:20,}},
	}
	store.discount(50)
	store.print()

	t := &toy{product{
		name:  "snake and ladder",
		price: 130,
	}}

	fmt.Println(t)

	//var(
	//	products = book{name:  "History of Iran", price: 1000,}
	//	Hitman = game{name:  "Hitman", price: 100,}
	//	tetris = game{name:  "Tetris", price: 50,}
	//	purepuz = puzzle{name:"purepuz", price: 30,}
	//	Mench = toy{name: "Mench", price:20,}
	//)
	//var store list
	//store = append(store, &Hitman, &tetris, products, purepuz, &Mench)
	//store.discount(50)
	//store.print()

	//var p printer
	//p = &tetris
	//p = products
	//p = &Hitman
	//
	//
	//fmt.Println(p)

	//p.(&Hitman).discount(30)
	//printbook(products)
	//printgame(gameproducts)
	//printgame(tetris)
	//products.print()
	//gameproducts.print()
	//tetris.print()
	////(&tetris).discount(50)
	//tetris.discount(50)
	//
	//var items []*game
	////items = [tetris, gameproducts]
	//lst := list(items)
	//lst.print()
}
