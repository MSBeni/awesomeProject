package main

import(
	"fmt"
)

func main(){
	arr := [...] int{1,2,3}
	arr2 := [...] int{1,2,3}
	inc_2(arr2)
	fmt.Println(arr2)
	inc_ptr_2(&arr2)
	fmt.Println(arr2)
	inr_arr := inc(arr)
	inrptr_arr := inc_ptr(&arr)
	fmt.Println(inr_arr,inrptr_arr)
}

func inc(arra [3]int)[3]int {
	for i := range arra {
		arra[i]++
	}
	return arra
}
func inc_ptr(arrptr *[3]int)*[3]int {
	addr:= &arrptr
	arr_addr := *addr
	for i := range arr_addr{
		arr_addr[i]++
	}
	return arr_addr
}
func inc_2(arra [3]int) {
	for i := range arra {
		arra[i]++
	}
}
func inc_ptr_2(arrptr *[3]int){
	addr:= &arrptr
	arr_addr := *addr
	for i := range arr_addr{
		arr_addr[i]++
	}
}