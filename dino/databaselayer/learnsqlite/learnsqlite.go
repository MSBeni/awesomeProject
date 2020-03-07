package main


import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type animal struct {
	id int
	animalType string
	nickname string
	zone int
	age int
}

func main(){
	log.Println("Connecting to the database (sqlite) ...")
	db, err := sql.Open("sqlite3", "dino.db")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()

	// create db and sample data
	//err = createdatabase(db)
	//if err!=nil {
	//	log.Fatal(err)
	//}
	//err = insertsampledata(db)
	//if err!=nil {
	//	log.Fatal(err)
	//}

	// GENERAL QUERY WITH ARGUMENTS
	rows, err := db.Query("SELECT * FROM animals WHERE age>$1", 5) // both $ and the ?
	if err!= nil{
		log.Fatal(err)
	}
	handlerows(rows, err)

	//QUERY A SINGLE ROW -- JUST RETURN THE FIRST ONE
	row := db.QueryRow("SELECT * FROM animals WHERE age>?", 5)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println(a)


	//insert a row
	result, err := db.Exec("Insert into animals (animal_type,nickname,zone,age) values ('Carnotaurus', 'Carno', $1, $2)", 3, 22)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())


	//update a row

	res, err := db.Exec("Update animals set age = $1, zone=$2 where id = $3", 16, 1, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}


func createdatabase(db *sql.DB)error{
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS 
							animals(id INTEGER PRIMARY KEY AUTOINCREMENT, 
							animal_type TEXT,
							nickname TEXT,
							zone INTEGER, 
							age INTEGER)`)
	return err
}

func insertsampledata(db *sql.DB)error{
	_, err := db.Exec(`insert INTO animals (animal_type,nickname,zone,age) 
                        VALUES ('Tyrannosaurus rex','rex', 1, 10),
                        ('Velociraptor', 'rapto', 2, 15)`)
	return err
}

//func handlerows(rows *sql.Rows, err error){
//	if err!= nil{
//		log.Fatal(err)
//	}
//	rows.Close()
//	animals := []animal{}
//	for rows.Next(){
//		a := animal{}
//		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
//		if err!= nil{
//			log.Println(err)
//			continue
//		}
//		if err:=rows.Err(); err!=nil{
//			log.Fatal(err)
//		}
//		animals = append(animals, a)
//	}
//	fmt.Println(animals)
//}


func handlerows(rows *sql.Rows, err error) {
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
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
}
