package main

import(
	"fmt"
	"sort"
	"strings"
	s "github.com/inancgumus/prettyslice"
)

func main(){
	Ini_sli := []float64{18,12,14,69,2,36}
	new_Ini_sli := append([]float64(nil),Ini_sli...)
	new_In := new_Ini_sli[:3]
	sort.Float64s(new_In)
	fmt.Println(Ini_sli)
	fmt.Println(new_In)
	m := copy(new_In, Ini_sli)
	fmt.Println(m, new_In)

	the_str_slice :=[]string{"aa","asdd","dferfd"}
	len_str := len(the_str_slice)
	up_le_slice := make([]string,len_str,len_str)
	for i,el:= range the_str_slice{
		//up_le_slice = append(up_le_slice, strings.ToUpper(el))
		up_le_slice[i] = strings.ToUpper(el)
		s.Show("Uppercase",up_le_slice)
	}

	a := [][]int{{1,2,5},{44,55,78,1,123}}
	a_new := make([][]int,0,5)
	a_new = append(a_new, []int{4,5,7,8,2})
	a_new = append(a_new, []int{4,8,2})
	s.Show("a: ",a)
	s.Show("a_new: ",a_new)
	//fmt.Println(up_le_slice)
	//s.Show("Uppercase",up_le_slice)
}