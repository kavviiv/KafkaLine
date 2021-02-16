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
	//CLUID := kafka.Consumer()
	dataKafka := kafka.Consumer()
	db := database.FetchData()

	//fmt.Println("Kafka Line Id", CLUID.LineID)

	client := &http.Client{}
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"), linebot.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Line bot client ERROR: ", err)
	}

	for _, em := range db {
		if em.BotStatus == nil {
		}

		if em.BotStatus != nil {
			for *em.BotStatus == "connect" && em.LineUID == dataKafka.LineID {
				if dataKafka == nil {
					continue
				}
				if dataKafka != nil {
					if _, err := bot.PushMessage(dataKafka.LineID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดเวลาตรวจสภาพรถแล้ว")).Do(); err != nil {
						log.Print(err)
						break
					}
					break
				}

			}
		}
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/linemessage", func(c echo.Context) error {
		//log.Println("here")
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			log.Fatal("Line bot client ERROR: ", err)
		}

		fmt.Println("db = ", db)
		for _, event := range events {
			if event.Type == linebot.EventTypeFollow {
				a := event.Source.UserID
				log.Println("user add bot")
				for _, g := range db {
					if g.LineUID == event.Source.UserID {
						status = "connect"
						log.Println("Equals")
						if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err == nil {
							//log.Print(err)
							dbc := database.DBCon()
							sqlStatement := `UPDATE test SET bot_status = $1 WHERE line_id = $2`
							_, err = dbc.Exec(sqlStatement, status, a)
							if err != nil {
								//w.WriteHeader(http.StatusBadRequest)
								panic(err)
							}
							break
						}

						if g.LineUID != event.Source.UserID {
							status = "disconnect"
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("You're not connect to our service")).Do(); err != nil {
								//log.Print(err)
								dbc := database.DBCon()
								sqlStatement := `UPDATE test SET bot_status = $1 WHERE line_id = $2`
								_, err = dbc.Exec(sqlStatement, status, event.Source.UserID)
								if err != nil {
									//w.WriteHeader(http.StatusBadRequest)
									panic(err)
								}
							}
							break
						}

					}

					log.Println("a=", a)

				}
				// for _, em := range db {
				// 	for em.BotStatus == "Connect" && em.LineUID == a {
				// 		dataKafka := kafka.Consumer()
				// 		if dataKafka == nil {
				// 			continue
				// 		}
				// 		if dataKafka != nil {
				// 			if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดเวลาตรวจสภาพรถแล้ว")).Do(); err != nil {
				// 				log.Print(err)
				// 				break
				// 			}
				// 			break
				// 		}

				// 	}
				// }
				//for  == "Connect" {
				// dataKafka := kafka.Consumer()

				// if dataKafka == nil {
				// 	continue
				// }

				// if dataKafka != nil {
				// 	if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดเวลาตรวจสภาพรถแล้ว")).Do(); err != nil {
				// 		log.Print(err)
				// 		break
				// 	}
				// 	break

				// }

				//	}

				if event.Type == linebot.EventTypeUnfollow {
					log.Println("user blcok bot")
					if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
						log.Print(err)
					}
				}

			}
			// for _, em := range db {
			// 	if em.BotStatus == nil {
			// 	}

			// 	if em.BotStatus != nil {
			// 		for *em.BotStatus == "connect" && em.LineUID == event.Source.UserID {

			// 			dataKafka := kafka.Consumer()
			// 			if dataKafka == nil {
			// 				continue
			// 			}
			// 			if dataKafka != nil {
			// 				if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดเวลาตรวจสภาพรถแล้ว")).Do(); err != nil {
			// 					log.Print(err)
			// 					break
			// 				}
			// 				break
			// 			}

			// 		}
			// 	}
			// }

			// if event.AccountLink.Nonce == "ok" {
			// 	//if event.AccountLink.Result == "ok" {
			// 	for true {
			// 		if dataKafka == nil {
			// 			continue
			// 		}
			// 		if dataKafka != nil {
			// 			if _, err := bot.PushMessage(dataKafka.LineID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดเวลาตรวจสภาพรถแล้ว")).Do(); err != nil {
			// 				log.Print(err)
			// 				break
			// 			}
			// 			break
			// 		}
			// 		//}

			// 	}

			// 	// for _, em := range db {
			// 	// 	if em.BotStatus == nil {
			// 	// 	}

			// 	// 	if em.BotStatus != nil {
			// 	// 		for *em.BotStatus == "connect" && em.LineUID == dataKafka.LineID {

			// 	// 			if dataKafka == nil {
			// 	// 				continue
			// 	// 			}
			// 	// 			if dataKafka != nil {
			// 	// 				if _, err := bot.PushMessage(dataKafka.LineID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดเวลาตรวจสภาพรถแล้ว")).Do(); err != nil {
			// 	// 					log.Print(err)
			// 	// 					break
			// 	// 				}
			// 	// 				break
			// 	// 			}

			// 	// 		}
			// 	// 	}
			// 	// }

			// }

		}
		//		dataKafka := kafka.Consumer()

		for _, em := range db {
			if em.BotStatus == nil {
			}

			if em.BotStatus != nil {
				for *em.BotStatus == "connect" && em.LineUID == dataKafka.LineID {

					if dataKafka == nil {
						continue
					}
					if dataKafka != nil {
						if _, err := bot.PushMessage(dataKafka.LineID, linebot.NewTextMessage("รถหมายเลขทะเบียน "+dataKafka.CarID+" ถึงกำหนดเวลาตรวจสภาพรถแล้ว")).Do(); err != nil {
							log.Print(err)
							break
						}
						break
					}

				}
			}
		}

		return c.String(http.StatusOK, "OK!")
	})

	e.Logger.Fatal(e.Start("127.0.0.1:9090"))

}
