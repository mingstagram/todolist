package main

// 애클리케이션의 진입점으로 사용
import (
	"backend/internal/app"
	"backend/internal/db"
	"log"
	"net/http"
)

func main() {

	// DB 연결
	database := db.Connect()
	defer database.Close()

	// 라우터 초기화
	router := app.InitRouter(database)

	// 서버 실행
	log.Println("Server running on http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", router)) 
}