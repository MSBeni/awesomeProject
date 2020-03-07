package main

import(
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
	s "github.com/inancgumus/prettyslice"
)

func main(){
	//var new_array [5] string
	var new_array [5] int
	//new_slice:= []string{"abc"}
	//
	//fmt.Println(new_array,new_slice )
	var new_slice []int
	var new_slice_2 []int
	max,_ := strconv.Atoi(os.Args[1])
	const_num := 5
	//for i:=0; i<const_num;{
	//	a := 0
	//	rand.Seed(time.Now().UnixNano())
	//	rand_num := rand.Intn(const_num)+1
	//	fmt.Println(rand_num)
	//	for _,nn:= range new_array{
	//		if rand_num == nn{
	//			a++
	//		}
	//		fmt.Println("array: ",nn)
	//	}
	//	if a==0{
	//		new_array[i] = rand_num
	//		i++
	//	}
	//
	//}
	Loop:
		for i:=0; i<const_num;{
			rand.Seed(time.Now().UnixNano())
			rand_num := rand.Intn(const_num)+1
			//fmt.Println(rand_num)
			for _,nn:= range new_array{
				if rand_num == nn{
					continue Loop
				}
				//fmt.Println("array: ",nn)
			}

			new_array[i] = rand_num
			i++
		}
	for _,el:=range new_array{
		new_slice_2 = append(new_slice_2,el)
	}
	fmt.Println(new_array)
	Loop2:
		for i:=0; i<max;{
			rand.Seed(time.Now().UnixNano())
			rand_num2 := rand.Intn(max)+1
			//fmt.Println(rand_num)
			for _,nn:= range new_slice{
				if rand_num2 == nn{
					continue Loop2
				}
				//fmt.Println("array: ",nn)
			}

			new_slice= append(new_slice,rand_num2)
			i++
		}
	fmt.Println(new_slice)
	new_slice = append(new_slice, new_slice_2...)
	fmt.Println(new_slice)
	sort.Ints(new_slice)
	fmt.Println(new_slice)
	s.Show("slice", new_slice)
}

