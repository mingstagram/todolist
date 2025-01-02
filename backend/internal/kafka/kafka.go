package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func InitKafkaProducer(broker string, topic string) {
	Writer = &kafka.Writer{
		Addr: kafka.TCP(broker),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func SendMessage(message string) error {
	err := Writer.WriteMessages(nil, kafka.Message {
		Value: []byte(message),
	})
	if err != nil {
		log.Println("Error while sending message: ", err)
		return err 
	}
	return nil
}

func CloseProducer() {
	if err := Writer.Close(); err != nil {
		log.Println("Failed to close Kafka writer: ", err)
	}
}