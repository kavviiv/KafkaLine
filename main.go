package main

import (
	database "KafkaLine/Database"
	kafka "KafkaLine/Kafka"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

// func contains(db []string, str string) bool {
// 	for _, v := range db {
// 		if v == str {
// 			return true
// 		}
// 	}

// 	return false

// }
var status string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.
			Fatal("Error loading .env file")
	}

	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	// a := kafka.Consumer()
	// log.Println(a)
	kafka := kafka.NewKafkaReader(kafkaHost, kafkaTopic)
	CLUID := kafka.Consumer()
	fmt.Println("Kafka Line Id", CLUID.LineID)

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
		//log.Println("here")
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			log.Fatal("Line bot client ERROR: ", err)
		}
		db := database.FetchData()
		fmt.Println("db = ", db)
		for _, event := range events {
			if event.Type == linebot.EventTypeFollow {
				a := event.Source.UserID
				log.Println("user add bot")
				for _, g := range db {
					if g.LineUID == event.Source.UserID {
						status = "connect"
						log.Println("Equals")
						if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
							log.Print(err)
							break
						}

						if g.LineUID != event.Source.UserID {
							if _, err := bot.PushMessage(a, linebot.NewTextMessage("You're not connect to our service")).Do(); err != nil {
								log.Print(err)
							}
							break
						}

					}

					log.Println("a=", a)

				}

				if event.Type == linebot.EventTypeUnfollow {
					log.Println("user blcok bot")
					if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
						log.Print(err)
					}
				}

			}

		}

		return c.String(http.StatusOK, "OK!")
	})

	e.Logger.Fatal(e.Start(":9090"))

}
