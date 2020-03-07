package main

import(
	"fmt"
)

type song struct{
	Album_name string
	songs_list []string
	price float64
}
func main(){
	type person struct{
		name, familyname, father string
		age int
		Iraniian bool

	}
	var Ali person
	Ali.name = "Ali"
	Ali.familyname = "Aminian"
	Ali.father = "Hashem"
	Ali.age = 32
	Ali.Iraniian = true

	fmt.Printf("\n Ali: %+v \n", Ali)
	fmt.Printf("\n Ali's last name: %#v \n", Ali.familyname)


	var MATRIX song
	MATRIX.Album_name = "Matrix"
	MATRIX.songs_list = []string{"love", "Money"}
	MATRIX.price = 129.3
	fmt.Printf("\n SONG: %+v \n", MATRIX)

}
