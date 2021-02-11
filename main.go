package main

import (
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main (
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ID := osGetenv("UID")
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

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				if _, err := bot.PushMessage(ID, linebot.NewTextMessage("hello")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
		return c.String(http.StatusOK, "OK!")
	})

	e.Logger.Fatal(e.Start(":9090"))
	
)


