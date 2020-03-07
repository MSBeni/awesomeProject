package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	// "postgres://pqgotest:password@localhost/pqgotest?sslmode=disable"
	//db, err := sql.Open("postgres", "postgres://dino:13ms13@localhost/dino?sslmode=disable")
	//"user=postgres dbname=dino sslmode=disable"
	db, err := sql.Open("postgres", "user=postgres password=13ms13 dbname=postgres sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// GENERAL QUERY WITH ARGUMENTS
	rows, err := db.Query("SELECT * FROM public.animals WHERE age>$1", 5)
	handlerows(rows, err)
	// QUERY A SINGLE ROW
	row := db.QueryRow("SELECT * FROM public.animals WHERE age>$1", 5)
	a:= animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(a)

	//insert a row
	//result, err := db.Exec("Insert into animals (animal_type,nickname,zone,age) values ('Carnotaurus', 'Carno', $1, $2)", 3, 22)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(result.LastInsertId()) //not supported here
	//fmt.Println(result.RowsAffected())


	//update a row
	//res, err := db.Exec("Update animals set age = $1 where id = $2", 16, 2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res.LastInsertId()) //not supported here
	//fmt.Println(res.RowsAffected())

	//var id int
	//db.QueryRow("Update animals set age = $1 where id = $2 returning id", 16, 2).Scan(&id)
	//fmt.Println("id returned:", id)
	fmt.Println("Statements...")
	stmt, err := db.Prepare("SELECT * FROM animals WHERE age>$1")
	if err != nil{
		log.Fatal(err)
	}
	defer stmt.Close()

	//let's try with age>5
	rows, err = stmt.Query(5)
	handlerows(rows, err)

	//let's try with age>10
	rows, err = stmt.Query(10)
	handlerows(rows, err)

	testTransaction(db)

}

func handlerows(rows *sql.Rows, err error){
	if err!=nil{
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next(){
		a:= animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err!=nil{
			log.Println(err)
			continue
		}
		if err:= rows.Err(); err!=nil{
			log.Fatal(err)
		}
		animals = append(animals, a)

	}
	fmt.Println(animals)

}

func testTransaction(db *sql.DB){
	fmt.Println("Transactions ...")
	tx, err := db.Begin()
	if err!=nil{
		log.Fatal(err)
	}
	//tx.Rollback()

	stmt, err := db.Prepare("SELECT * FROM animals WHERE age>$1")
	if err != nil{
		log.Fatal(err)
	}
	defer stmt.Close()
	//let's try with age>5
	rows, err := stmt.Query(15)
	handlerows(rows, err)

	//let's try with age>10
	rows, err = stmt.Query(17)
	handlerows(rows, err)

	result, err := tx.Exec("Update animals set age = $1 where id = $2", 14, 2)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
	err = tx.Commit()
	if err != nil{
		log.Fatal(err)
	}
}