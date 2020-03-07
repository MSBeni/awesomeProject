package main

import(
	"encoding/json"
	"fmt"
)
type permissions map[string]bool

type users struct{
	Username string
	Pass string //`json:"-"`
	Accesses permissions `json:"Permissions,omitempty"`   // encoding the fields - the omitempty field prevent the nill part to be shown

}

func main(){
	web_users:= []users{
		{Username: "Ali", Pass:     "123", Accesses: nil,},
		{Username: "Noushi", Pass:     "1234", Accesses: nil,},
		{Username: "Mo", Pass:     "1235", Accesses: permissions{"web":true,"domains":false},},
	}
	// ENCODING THE JSON DATA
	//out, err := json.Marshal(web_users)
	out, err := json.MarshalIndent(web_users, "", "\t")
	if err!=nil{
		fmt.Println("Error!!!")
		return
	}
	fmt.Println(string(out))  // IT OUGHT TO BE string(out)
}