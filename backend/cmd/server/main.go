package main

// 애클리케이션의 진입점으로 사용
import (
	"backend/internal/app"
	"backend/internal/db"
	_ "backend/internal/kafka"
	"log"
	"net/http"
)

func main() { 
	// DB 연결
	database := db.Connect()
	defer database.Close()

	// // Kafka Producer 초기화
	// kafka.InitKafkaProducer("localhost:9092", "task-updates")

	// // Kafka Consumer는 별도의 Go 루틴으로 실행 (실시간 알림을 위한 WebSocket 등으로 처리 가능)
	// go kafka.ConsumeMessages("localhost:9092", "task-updates")

	// 라우터 초기화
	router := app.InitRouter(database)

	// 서버 실행
	log.Println("Server running on http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", router)) 
}