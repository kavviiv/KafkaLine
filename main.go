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

var status, UserConfirm string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.
			Fatal("Error loading .env file")
	}
	database.FetchData()

	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafka := kafka.NewKafkaReader(kafkaHost, kafkaTopic)
	dataKafka := kafka.Consumer()
	db := database.FetchData()
	car := dataKafka.CarID
	//x := dataKafka.CarID

	client := &http.Client{}
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"), linebot.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Line bot client ERROR: ", err)
	}

	var contents []linebot.FlexComponent
	var head []linebot.FlexComponent
	var foot []linebot.FlexComponent
	//	var action []linebot.Action
	text := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Align:  "center",
		Text:   "รถหมายเลขทะเบียน",
		Weight: "regular",
		Size:   linebot.FlexTextSizeTypeMd,
	}

	text2 := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Align:  "center",
		Text:   car,
		Weight: "bold",
		Size:   linebot.FlexTextSizeTypeXl,
	}
	text3 := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Align:  "center",
		Text:   "ถึงกำหนดตรวจสภาพรถแล้ว",
		Weight: "regular",
		Size:   linebot.FlexTextSizeTypeMd,
	}

	htext := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Align:  "center",
		Text:   "ตรวจสภาพรถ",
		Weight: "bold",
		Size:   linebot.FlexTextSizeTypeXl,
	}

	foot1 := linebot.ButtonComponent{
		Type: "button",
		Action: &linebot.PostbackAction{
			Label: "ตกลง",
			Data:  "confirm",
			//Text:        "ใช่",
			DisplayText: "ยืนยัน",
		},
		Gravity: "center",
	}

	contents = append(contents, &text, &text2, &text3)
	head = append(head, &htext)
	foot = append(foot, &foot1)

	footer := linebot.BoxComponent{
		Type:     "box",
		Layout:   "vertical",
		Contents: foot,
	}
	header := linebot.BoxComponent{
		Type:     "box",
		Layout:   "vertical",
		Contents: head,
	}
	body := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: contents,
	}

	bubble := linebot.BubbleContainer{
		Type:   linebot.FlexContainerTypeBubble,
		Body:   &body,
		Header: &header,
		Footer: &footer,
	}
	flexMessage := linebot.NewFlexMessage("ครบกำหนดตรวจสภาพรถ", &bubble)

	for _, em := range db {
		if em.BotStatus == nil {
		}

		//for true {
		if em.BotStatus != nil {
			for *em.BotStatus == "connect" && em.LineUID == dataKafka.LineID {
				//for true {
				if dataKafka == nil {
					continue
				}

				if dataKafka != nil {
					if _, err := bot.PushMessage(dataKafka.LineID, flexMessage).Do(); err != nil {
						log.Print(err)
						break
					}
					break
				}

			}
		}

	}

	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
	{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "กำลังตรวจสถานะ",
			  "weight": "bold",
			  "align": "center",
			  "contents": []
			}
		  ]
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "image",
			  "url": "https://www.flaticon.com/free-icon/sand-clock_889843?term=wait&page=1&position=5&page=1&position=5&related_id=889843&origin=search",
			  "size": "md"
			},
			{
			  "type": "text",
			  "text": "ระบบกำลังตรวจสถานะ",
			  "align": "center",
			  "contents": []
			},
			{
			  "type": "text",
			  "text": "การเข้าตรวจสภาพรถของท่าน",
			  "align": "center",
			  "contents": []
			}
		  ]
		}
	  }
	
	`))
	if err != nil {
		log.Println(err)
	}
	// New Flex Message
	waitForConfirm := linebot.NewFlexMessage("FlexWithJSON", flexContainer)

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
						log.Println(status)
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
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("You're not connect to our service")).Do(); err != nil {
								//log.Print(err)

							}
							break
						}

					}

					log.Println("a=", a)

				}

			}
			if event.Type == linebot.EventTypeUnfollow {
				status = "disconnect"
				log.Println(status)
				log.Println("user blcok bot")
				dbc := database.DBCon()
				sqlStatement := `UPDATE test SET bot_status = $1 WHERE line_id = $2`
				_, err = dbc.Exec(sqlStatement, status, event.Source.UserID)
				if err != nil {
					//w.WriteHeader(http.StatusBadRequest)
					panic(err)
				}

				if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Bye bitch")).Do(); err != nil {
					log.Print(err)

				}
			}

			if event.Type == linebot.EventTypePostback {
				UserConfirm = "Confirm"
				log.Println(UserConfirm)
				if _, err := bot.PushMessage(event.Source.UserID, waitForConfirm).Do(); err != nil {
					log.Print(err)

				}

			}

		}

		return c.String(http.StatusOK, "OK!")
	})

	e.Logger.Fatal(e.Start("127.0.0.1:9090"))

}
