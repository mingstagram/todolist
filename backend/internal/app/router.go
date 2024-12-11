package app

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"
	"database/sql"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func setupCORS(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for all responses
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type") // Allow specific headers
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allow specific methods
}

func InitRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// CORS 미들웨어 추가
	router.Use(middleware.CORS)

	// 라우트 설정
	// tasks
	tasksRepository := repositories.NewTasksRepository(db)
	tasksService := services.NewTasksService(tasksRepository)
	tasksHandler := handlers.NewTasksHandler(tasksService)

	// users
	usersRepository := repositories.NewUsersRepository(db)

	// auth
	authService := services.NewAuthService(usersRepository)
	authHandler := handlers.NewAuthHandler(authService)

	// Auth routes
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/auth/signup", authHandler.SaveUsers).Methods("POST")

	// 보호된 라우트 (JWT 인증 필요)
	protectedRoutes := router.PathPrefix("/tasks").Subrouter()
	protectedRoutes.Use(middleware.JWTMiddleware)

	// CORS preflight (OPTIONS) 요청을 모든 "/tasks" 엔드포인트에 대해 처리
	protectedRoutes.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			setupCORS(w, r) // CORS 헤더 설정
			w.WriteHeader(http.StatusOK)
			return
		}
	}).Methods(http.MethodOptions)

	// Other task routes
	// protectedRoutes.HandleFunc("/checked", tasksHandler.UpdateChecked).Methods("PUT")
	// protectedRoutes.HandleFunc("/count", tasksHandler.CountTasks).Methods("GET")
	protectedRoutes.HandleFunc("", tasksHandler.GetTasksForDate).Methods("GET")
	protectedRoutes.HandleFunc("", tasksHandler.SaveTasks).Methods("POST")
	protectedRoutes.HandleFunc("", tasksHandler.DeleteTasks).Methods("DELETE")

	checkedRoutes := router.PathPrefix("/tasks/checked").Subrouter()
	checkedRoutes.Use(middleware.JWTMiddleware)

	// CORS preflight (OPTIONS) 요청을 모든 "/tasks" 엔드포인트에 대해 처리
	checkedRoutes.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			setupCORS(w, r) // CORS 헤더 설정
			w.WriteHeader(http.StatusOK)
			return
		}
	}).Methods(http.MethodOptions)
 
	checkedRoutes.HandleFunc("", tasksHandler.UpdateChecked).Methods("PUT")

	countRoutes := router.PathPrefix("/tasks/count").Subrouter()
	countRoutes.Use(middleware.JWTMiddleware)

	// CORS preflight (OPTIONS) 요청을 모든 "/tasks" 엔드포인트에 대해 처리
	countRoutes.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			setupCORS(w, r) // CORS 헤더 설정
			w.WriteHeader(http.StatusOK)
			return
		}
	}).Methods(http.MethodOptions)


	countRoutes.HandleFunc("", tasksHandler.CountTasks).Methods("GET")

	return router
}