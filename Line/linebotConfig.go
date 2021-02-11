package line

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func linebotConfig() *linebot.Client {
	err := godotenv.Load()
	log.Printf("%s", err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}
