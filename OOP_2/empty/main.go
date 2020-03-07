package main

import "fmt"

func main(){
	var any interface{}
	any = []int{1,2,3}
	any = map[int]bool{1:true,2:false,3:true}
	any = 3
	any = any.(int)*3
	fmt.Println(any)
}
