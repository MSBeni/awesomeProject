package main

import(
	"fmt"
	"os"
	"awesomeProject/mapfiles/packages"
	"awesomeProject/mapfiles/scanbyword"
)

func main(){
	dict := map[string]string{}
	n_dict_2 := make(map[string]string)
	dict["a"] = "Hello"
	dict["b"] = "HeHe"

	new_dict := dict
	new_dict["c"] = "New Hi"

	for k,v:= range dict{
		n_dict_2[k]=v
	}
	n_dict_2["d"] = "Very New Hi"
	//fmt.Println(dict["a"])
	query := os.Args[1]
	//value, ok := dict[query]
	if value, ok := dict[query]; ok {
		fmt.Println(dict[query])
		fmt.Println(value)
	}else{
		fmt.Printf(" value is not available \n")
	}
	fmt.Printf(" %q \n", dict)
	fmt.Printf(" %q \n", new_dict)
	fmt.Printf(" %q \n", n_dict_2)
	scanbyword.Scanbyword()
	small_pro.Small_pro()

}