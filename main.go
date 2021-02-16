package main

import (
	database "KafkaLine/Database"
	kafka "KafkaLine/Kafka"

	//kafka "KafkaLine/Kafka"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

var status string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.
			Fatal("Error loading .env file")
	}
	database.FetchData()

	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	// a := kafka.Consumer()
	// log.Println(a)
	kafka := kafka.NewKafkaReader(kafkaHost, kafkaTopic)
	CLUID := kafka.Consumer()
	fmt.Println("Kafka Line Id", CLUID.LineID)

	client := &http.Client{}
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"), linebot.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Line bot client ERROR: ", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/linemessage", func(c echo.Context) error {
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			log.Fatal("Line bot client ERROR: ", err)
		}
		db := database.FetchData()
		fmt.Println("db = ", db)
		for _, event := range events {
			log.Println("start Event")
			if event.Type == linebot.EventTypeFollow {
				log.Println("user add bot")
				for _, g := range db {
					if g.LineUID == event.Source.UserID {
						status = "connect"
						log.Println("Equals")
						if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
							log.Print(err)
							break
						}
						// for i := 0; i <= 4; i++ {

						// 	dataKafka := kafka.Consumer()

						// 	if dataKafka == nil {
						// 		fmt.Println("Null")
						// 		continue
						// 	}
						// 	if i == 0 {
						// 		log.Println("Kafka 0 =", dataKafka)
						// 		log.Println(i)
						// 		if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(" bitch")).Do(); err != nil {
						// 			log.Print(err)
						// 		}
						// 	}
						// 	if i == 2 {
						// 		log.Println("Kafka 2 =", dataKafka)
						// 		log.Println(i)
						// 		if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("ye bitch")).Do(); err != nil {
						// 			log.Print(err)
						// 		}
						// 	}
						// 	if i == 4 {
						// 		log.Println("Kafka 4 =", dataKafka)
						// 		log.Println(i)
						// 		if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(" bitch")).Do(); err != nil {
						// 			log.Print(err)
						// 		}
						// 	}
						// 	log.Println("index ", i)
						// 	log.Println("------------------------------------------------------------------")
						// }

						if g.LineUID != event.Source.UserID {
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("You're not connect to our service")).Do(); err != nil {
								log.Print(err)
							}
							break
						}

					}

					log.Println("Follower UserID = ", event.Source.UserID)

				}

				if event.Type == linebot.EventTypeUnfollow {
					log.Println("user blcok bot")
					if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
						log.Print(err)
					}
				}

			}
			// //sendMessages()
			// if event.Type == linebot.EventType(linebot.AccountLinkResultOK) {

			// 	for i := 0; i <= 4; i++ {

			// 		dataKafka := kafka.Consumer()

			// 		if dataKafka == nil {
			// 			fmt.Println("Null")
			// 			continue
			// 		}
			// 		if i == 0 {
			// 			log.Println("Kafka 0 =", dataKafka)
			// 			log.Println(i)
			// 			if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
			// 				log.Print(err)
			// 			}
			// 		}
			// 		if i == 2 {
			// 			log.Println("Kafka 2 =", dataKafka)
			// 			log.Println(i)
			// 			if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
			// 				log.Print(err)
			// 			}
			// 		}
			// 		if i == 4 {
			// 			log.Println("Kafka 4 =", dataKafka)
			// 			log.Println(i)
			// 			if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
			// 				log.Print(err)
			// 			}
			// 		}
			// 		log.Println("index ", i)
			// 		log.Println("------------------------------------------------------------------")
			// 	}
			// 	// if linebot.AccountLinkResult == linebot.AccountLinkResultOK {

			// 	// }

			// }

		}
		if status == "connect" {
			for i := 0; i <= 4; i++ {

				dataKafka := kafka.Consumer()

				if dataKafka == nil {
					fmt.Println("Null")
					continue
				}
				if i == 0 {
					log.Println("Kafka 0 =", dataKafka)
					log.Println(i)
					if _, err := bot.PushMessage("U357dfdc149b28a464a83819e7fedd332", linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
						log.Print(err)
					}
				}
				if i == 2 {
					log.Println("Kafka 2 =", dataKafka)
					log.Println(i)
					if _, err := bot.PushMessage("U357dfdc149b28a464a83819e7fedd332", linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
						log.Print(err)
					}
				}
				if i == 4 {
					log.Println("Kafka 4 =", dataKafka)
					log.Println(i)
					if _, err := bot.PushMessage("U357dfdc149b28a464a83819e7fedd332", linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
						log.Print(err)
					}
				}
				log.Println("index ", i)
				log.Println("------------------------------------------------------------------")
			}
		}
		// for linebot.AccountLinkResultOK == "OK" {
		// 	log.Println("hereh")

		// 	for {
		// 		dataKafka := kafka.Consumer()
		// 		if dataKafka == nil {
		// 			fmt.Println("Null")
		// 			continue
		// 		}
		// 		if dataKafka != nil {
		// 			log.Println("4556")
		// 			if _, err := bot.PushMessage("U357dfdc149b28a464a83819e7fedd332", linebot.NewTextMessage(dataKafka.CarID)).Do(); err != nil {
		// 				log.Print(err)
		// 			}
		// 			break

		// 		}
		// 	}

		// }

		return c.String(http.StatusOK, "OK!")
	})

	e.Logger.Fatal(e.Start(":9090"))

}
