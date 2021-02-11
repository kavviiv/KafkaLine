package database

import (
	"database/sql"
	"fmt"
	"log"
)

// DBCon ;

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

func FetchData() []LineID {
	db := DBCon()
	source := LineID{}
	var data []LineID

	// WHERE return user_id
	rows, err := db.Query("SELECT  line_id, user_id FROM test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&source.LineUID, &source.UID)

		if err != nil {
			fmt.Print(err)
		}

		data = append(data, source)
	}

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", source)
	defer db.Close()

	return data
}
