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
	rows, err := db.Query("SELECT  line_id, bot_status FROM test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&source.LineUID, &source.BotStatus)
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

func FetchData2() []Data {
	db := DBCon()
	defer db.Close()

	var data []Data
	dataList := Data{}

	// WHERE return data at user_id
	rows, err := db.Query("SELECT user_id, line_id, car_id FROM test")
	if err != nil {
		log.Fatalln("Error database query in Fetch Data")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&dataList.UserID, &dataList.LineID, &dataList.CarID)
		if err != nil {
			fmt.Println("Error row scan in Fetch Data")
			fmt.Print(err)
		}
		data = append(data, dataList)
	}
	return data
}
