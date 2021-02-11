package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := Database.FetchData()

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

		for _, event := range events {
			if event.Type == linebot.EventTypeFollow {

				a := event.Source.UserID
				log.Println(a)

				log.Println("user add bot")

				if _, err := bot.PushMessage(a, linebot.NewTextMessage("Welcome to our service")).Do(); err != nil {
					log.Print(err)
				}

			}

		}
		return c.String(http.StatusOK, "OK!")
	})

	e.Logger.Fatal(e.Start(":9090"))

}
