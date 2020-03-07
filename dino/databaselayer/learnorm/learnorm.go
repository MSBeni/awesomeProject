package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"fmt"
)

type animal struct {
	//gorm.Model
	ID         int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	Animaltype string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zone       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

func main(){
	db, err := gorm.Open("sqlite3", "dino.db")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	// GORM TAKE THE NAME OF THE OBJECT AND ADD THE s TO THE NAME TO BUILD THE TABLE

	db.DropTableIfExists(&animal{})
	db.Table("dinos").DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{})   // will add any missing fields, will add 's' to the struct name

	// CREATE THE TABLE WITH THE SPECIFIC NAME
	db.Table("dinos").CreateTable(&animal{})
	db.Table("dinos").AutoMigrate(&animal{})


	//inserts:
	a := animal{
		Animaltype: "Tyrannosaurus rex",
		Nickname:   "rex",
		Zone:       1,
		Age:        11,
	}
	//db.Save(&a)
	db.Create(&a)
	db.Table("dinos").Create(&a)
	a = animal{
		Animaltype: "Velociraptor",
		Nickname:   "rapto",
		Zone:       2,
		Age:        15,
	}
	//db.Save(&a)
	db.Save(&a)
	//queries
	animals := []animal{}
	db.Table("dinos").Find(&animals, "age > ?", 10) //first
	fmt.Println(animals)

	animals = []animal{}
	db.Find(&animals, "age > ?", 12) //first
	fmt.Println(animals)

	//updates
	//db.Table("animals").Where("nickname = ? and zone= ?", "rapto", 2).Update("age", 16)
}