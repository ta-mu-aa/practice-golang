package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age int
}

func main(){
	Dbconnection, _ := sql.Open("sqlite3", "./example.sql")
	defer Dbconnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person (
						name STRING,
						age INT)`
	_, err := Dbconnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = Dbconnection.Exec(cmd, "Nancy",20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = Dbconnection.Exec(cmd, 25, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

// 	cmd = "SELECT * FROM person"
// 	rows, _ := Dbconnection.Query(cmd)
// 	defer rows.Close()
// 	var pp []Person
// 	for rows.Next(){
// 		var p Person
// 		err := rows.Scan(&p.Name, &p.Age)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		pp = append(pp, p)
// 	}
// 	for _, p := range pp{
// 		fmt.Println(p.Name, p.Age)
// 	}

	// cmd = "SELECT * FROM person where age = ?"
	// row := Dbconnection.QueryRow(cmd, 20)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows{
	// 		log.Println("No row")
	// 	}else{
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	cmd = "DELETE FROM person where name = ?"
	_, err = Dbconnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}
