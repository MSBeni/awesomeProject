package main

type toy struct {
	product
}
//type toy struct {
//	name string
//	price money
//}
// now we are making methods
//func (t *toy)print(){
//	fmt.Printf("%-15s: %s\n", t.name, t.price.string())
//}
//
//func (t *toy)discount(ratio float64){
//	t.price*=money((100-ratio)/100)
//	fmt.Printf("The price of the %s after discount: %.2f\n", t.name, t.price)
//}