package app

import (
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	
	tasksRepository := repositories.NewTasksRepository(db)
	tasksService := services.NewTasksService(tasksRepository)
	tasksHandler := handlers.NewTasksHandler(tasksService)

	router.HandleFunc("/tasks/today", tasksHandler.GetTodayTasks).Methods("GET")
	router.HandleFunc("/tasks/date", tasksHandler.GetTasksForDate).Methods("GET")
	router.HandleFunc("/tasks", tasksHandler.SaveTasks).Methods("POST")
 
	return router
}
