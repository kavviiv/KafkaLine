package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var dbdata []UserLine

//DBCon :
func DBCon() *sql.DB {
	conndb := "user=postgres dbname=postgres password=130242 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", conndb)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	fmt.Println("_________________________________________")

	return db
}

// FetchData :
func FetchData() []UserLine {
	db := DBCon()
	source := UserLine{}
	var data []UserLine

	// WHERE return user_id
	rows, err := db.Query("SELECT  line_id FROM test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&source.LineUID)
		if err != nil {
			fmt.Print(err)
		}
		data = append(data, source)
	}

	log.Printf("555 %+v\n", data)
	//dbdata := data
	//fmt.Printf("666 %+v\n", source)
	defer db.Close()

	return data
}
