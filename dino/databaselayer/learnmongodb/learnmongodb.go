package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"log"
)

type animal struct {
	//ID         int    `bson:"id"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

func main(){
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err!=nil{
		log.Fatal(err)
	}
	defer session.Close()

	// GET A COLLECTION
	animalcollection := session.DB("Dino").C("animals")
	//animals := []interface{}{animal{
	//	AnimalType: "Tyrannosaurus rex",
	//	Nickname:   "rex",
	//	Zone:       1,
	//	Age:        11,
	//}, animal{
	//	AnimalType: "Velociraptor",
	//	Nickname:   "rapto",
	//	Zone:       2,
	//	Age:        17,
	//}, animal{
	//	AnimalType: "Velociraptor",
	//	Nickname:   "Velo",
	//	Zone:       2,
	//	Age:        9,
	//},
	//}
	//err = animalcollection.Insert(animals...)
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//
	//err = animalcollection.Update(bson.M{"nickname":"rapto"}, bson.M{"$set":bson.M{"age":"19"}})
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//
	//err = animalcollection.Remove(bson.M{"nickname":"rapto"})
	//if err!=nil{
	//	log.Fatal(err)
	//}

	query := bson.M{
		"age": bson.M{
			"$gt": 5,
		},
		"zone": bson.M{
			"$in": []int{1, 2},
		},
	}
	results := []animal{}
	animalcollection.Find(query).All(&results) //.one
	fmt.Println(results)

	//one result
	result := animal{}
	animalcollection.Find(query).One(&result)
	fmt.Println(result)

}
