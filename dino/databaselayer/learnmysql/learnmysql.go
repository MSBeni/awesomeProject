package learnmysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	// connect to Database
	db, err := sql.Open("mysql", "root:Password@/Dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// general query with arguments
	rows, err := db.Query("SELECT * FROM Dino.animals WHERE age>?", 10)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		animals = append(animals, a)

	}
	fmt.Println(animals)

	// INSERT A ROW
	//result, err := db.Exec("INSERT INTO Dino.animals (animal_type,nickname,zone,age) VALUES ('Carnotaurus', 'carno' ,3,22)")
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())

	// UPDATE A ROW
	age := 14
	id := 2
	result, err := db.Exec("UPDATE Dino.animals SET age=? WHERE id=?", age, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
