package app

import (
	"backend/internal/handlers"
	"backend/internal/services"
	"database/sql"

	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// 서비스 계층 초기화
	testboardService := services.NewTestboardService(db)

	// 핸들러 초기화
	testboardHandler := handlers.NewTestboardHandler(testboardService)
 
	// 라우팅 설정
	// router.HandleFunc("/testboards", testboardHandler.GetTestboards).Methods("GET")
    // router.HandleFunc("/testboards/{id}", testboardHandler.GetTestboard).Methods("GET")
    // router.HandleFunc("/testboards", testboardHandler.CreateTestboard).Methods("POST")
    // router.HandleFunc("/testboards/{id}", testboardHandler.UpdateTestboard).Methods("PUT")
    // router.HandleFunc("/testboards/{id}", testboardHandler.DeleteTestboard).Methods("DELETE")
	router.HandleFunc("/testboards", testboardHandler.GetAllTestboards).Methods("GET")

	return router
}
