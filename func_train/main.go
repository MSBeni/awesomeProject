package main

import(
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main(){
	data_rec := read_data()
	jsonfile, err := json_conv(data_rec)
	fmt.Printf("the json value is: %s and the error is %s \n", jsonfile, err)

}

func read_data()[]byte{
	var input []byte
	for in :=bufio.NewScanner(os.Stdin); in.Scan();{
		input = append(input, in.Bytes()...)
		input = append(input, '\n')
	}
	return input
}

func json_conv(input []byte)(interface{}, error){
	var courses map[string]interface{}
	err := json.Unmarshal(input, &courses)

	if err!= nil{
		fmt.Printf("Error!!!! %s", err)
	}
	//for _,val:=range courses{
	//	return(val)
	//}

	return courses, err
}
