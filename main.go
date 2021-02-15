package main

import (
	database "KafkaLine/Database"
	kafka "KafkaLine/Kafka"
	"bytes"
	"encoding/json"
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
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			log.Fatal("Line bot client ERROR: ", err)
		}

		//var tata []string
		///tata = append(tata)
		//var data []database.UserLine
		//	CLUID := kafka.Consumer()
		db := database.FetchData()
		fmt.Println("db = ", db)
		for _, event := range events {
			if linebot.AccountLinkResult(CLUID.LineID) == linebot.AccountLinkResultOK {
				log.Println("Account linked")
			}

			if event.Type == linebot.EventTypeFollow {
				a := event.Source.UserID
				log.Println("user add bot")
				for _, g := range db {

					if g.LineUID == event.Source.UserID && g.LineUID == a {
						status = "connect"
						log.Println(status)
						log.Println("Equals")
						if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
							log.Print(err)
							break
						}
						for true {
							dataKafka := kafka.Consumer()
							if dataKafka == nil {
								fmt.Println("Null")
								continue
							}
							if dataKafka != nil {
								if _, err := bot.PushMessage(CLUID.LineID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดตรวจสภาพรถแล้ว")).Do(); err != nil {
									log.Print(err)
								}
								break
							}
						}

						if g.LineUID != event.Source.UserID {
							if _, err := bot.PushMessage(a, linebot.NewTextMessage("You're not connect to our service")).Do(); err != nil {
								log.Print(err)
							}
							break
						}

					}

					type Payload struct {
						To string `json:"to"`
					}
					type Messages struct {
						Type string `json:"type"`
						Text string `json:"text"`
					}

					data := Payload{
						CLUID.CarID,
						// fill struct
					}
					payloadBytes, err := json.Marshal(data)
					if err != nil {
						// handle err
					}
					body := bytes.NewReader(payloadBytes)

					req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", body)
					if err != nil {
						// handle err
					}
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("Authorization", "Bearer"+os.Getenv("CHANNEL_TOKEN"))

					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						// handle err
					}
					defer resp.Body.Close()

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
