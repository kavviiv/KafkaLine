package main

//	"github.com/confluentinc/confluent-kafka-go/kafka"

func sendMessages() {
	// kafkaHost := os.Getenv("KAFKA_HOST")
	// kafkaTopic := os.Getenv("KAFKA_TOPIC")
	// bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"), linebot.WithHTTPClient(client))
	// if err != nil {
	// 	log.Fatal("Line bot client ERROR: ", err)

	// // a := kafka.Consumer()
	// // log.Println(a)
	// kafka := kafka.NewKafkaReader(kafkaHost, kafkaTopic)
	// CLUID := kafka.Consumer()
	// fmt.Println("Kafka Line Id", CLUID.LineID)

	// for true {
	// 	dataKafka := kafka.Consumer()
	// 	if dataKafka == nil {
	// 		fmt.Println("Null")
	// 		continue
	// 	}
	// 	if dataKafka != nil {
	// 		log.Println("4556")
	// 		if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(dataKafka.CarID)).Do(); err != nil {
	// 			log.Print(err)
	// 		}
	// 		continue

	// 	}
	// }
}
