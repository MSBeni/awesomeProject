package databaselayer

import (
	"awesomeProject/dino/databaselayer/dynowebportal"
	"encoding/json"
	"log"
	"os"
)
type configuration struct {
	ServerAddress string `json:"webserver"`
}

//// Creating a type to host the json data -- be careful the field of the struct which going to have the json data should
//// exportable ( upper case )
func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	// new make a pointer to the type passed to it
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting web server on address ", config.ServerAddress)
	dynowebportal.RunWebPortal(config.ServerAddress)
}
