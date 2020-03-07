package main

import(
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Ques struct{
	Question string
	Options []string
	Answer string
}


//type Course struct{
//	Courses Ques
//}
//type quiz struct{
//	Quizes course
//}
func main(){
	var input []byte
	for in :=bufio.NewScanner(os.Stdin); in.Scan();{
		input = append(input, in.Bytes()...)
		input = append(input, '\n')
	}
	//fmt.Println(string(input))
	//var quizes []quiz
	//var questions Ques
	var courses map[string]interface{}

	err := json.Unmarshal(input, &courses)

	if err!= nil{
		fmt.Println("Error!!!!")
		return
	}
	for _,val:=range courses{
		fmt.Println(val)
	}

}
