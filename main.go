package main

import (
	database "KafkaLine/Database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func contains(db []string, str string) bool {
	for _, v := range db {
		if v == str {
			return true
		}
	}

	return false

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.FetchData()

	client := &http.Client{}
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"), linebot.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Line bot client ERROR: ", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})

	//ID := "Uf6c3aa4a9460f9438f2ab25aea33ab67"

	e.POST("/linemessage", func(c echo.Context) error {
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			log.Fatal("Line bot client ERROR: ", err)
		}
		//var tata []string

		///tata = append(tata)

		//var data []database.UserLine
		db := database.FetchData()
		fmt.Println("db = ", db)
		for _, event := range events {
			if event.Type == linebot.EventTypeFollow {
				a := event.Source.UserID
				log.Println("user add bot")
				for _, g := range db {

					if g.LineUID == event.Source.UserID && g.LineUID == a {
						log.Println("Equals")
						//linebot.NewTextMessage("555")
						if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
							log.Print(err)
						}

					}

					if g.LineUID != event.Source.UserID {
						log.Println("Not Equals")
						if _, err := bot.PushMessage(a, linebot.NewTextMessage("You're not connect to our service")).Do(); err != nil {
							log.Print(err)
						}

					}
					// if g.LineUID == string(a) {

					// }
					// if g.LineUID != string(a) {
					// 	if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
					// 		log.Print(err)
					// 	}
					// }
				}

				// for _, g := range data {
				// 	if g.LineUID == a {
				// 		if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
				// 			log.Print(err)
				// 		}
				// 	}

				// 	if g.LineUID != a {
				// 		if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
				// 			log.Print(err)
				// 		}
				// 	}

				//	}

				log.Println("a=", a)
				//	log.Fatalln(db)
				//log.Println("X =", data)

				// if event.Type == linebot.EventTypeUnfollow {
				// 	log.Println("user blcok bot")
				// 	a := event.Source.UserID
				// 	if _, err := bot.PushMessage(a, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
				// 		log.Print(err)
				// 	}
				// }

			}

			if event.Type == linebot.EventTypeUnfollow {
				log.Println("user blcok bot")

				if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
					log.Print(err)
				}
			}

		}
		return c.String(http.StatusOK, "OK!")
	})

	e.Logger.Fatal(e.Start(":9090"))

}
