package main


func main(){
	var(
		Call_for_duty = game{name:"Call for duty 2020", price:1000, countries: map[string]bool{"USA":true, "CANADA": true, "North Korea":false},}
		FiFa = game{name:"FiFa 2020", price:500, countries: map[string]bool{"USA":true, "CANADA": true, "North Korea":false},}
		the_book = Book{name:"Kafka on the shore", price: 100,}
	)
	var items []*game
	items = append(items, &FiFa, &Call_for_duty)

	//Call_for_duty.print()
	//the_book.print()

	//(&Call_for_duty).discount(0.5)  // Go to this by itself
	Call_for_duty.discount(0.5)
	game.print(Call_for_duty)
	Book.print(the_book)
	game.print(FiFa)

	my := List(items)
	my.game_lst()


}
