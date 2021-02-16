package kafka

import (
	types "KafkaLine/types"

	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// KafkaReader :
type KafkaReader struct {
	reader *kafka.Reader
	writer *kafka.Writer
}

// NewKafkaReader :
func NewKafkaReader(address string, topic string) *KafkaReader {

	// อ่านที่ topic 1
	config := kafka.ReaderConfig{
		Brokers:     []string{address},
		Topic:       "KafkaTest",
		GroupID:     "a4",
		MaxBytes:    1e3,
		MinBytes:    1e2,
		StartOffset: kafka.LastOffset,
	}

	reader := kafka.NewReader(config)

	// เขียนที่ topic 2
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{address},
		Topic:   "KafkaTest",
	})

	fmt.Println("success read")

	return &KafkaReader{reader: reader, writer: writer}
}

// Consumer :
func (k *KafkaReader) Consumer() *types.DataConsumer {
	for {
		dataKafka := types.DataConsumer{}
		dataKafkaRead, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			return nil
		}
		if err := json.Unmarshal(dataKafkaRead.Value, &dataKafka); err != nil {
			return nil
		}
		fmt.Println("Kafka =", dataKafka)

		return &dataKafka
	}
}
