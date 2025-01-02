package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(broker string, topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic: topic,
		GroupID: "notification-group",
	})

	for {
		msg, err := reader.ReadMessage(nil)
		if err != nil {
			log.Println("Error while reading message: ", err)
			continue
		}

		// 여기서 WebSocket을 사용하여 클라이언트에 메시지를 전송하는 로직 추가
		log.Printf("Received message: %s", string(msg.Value))
	}
}