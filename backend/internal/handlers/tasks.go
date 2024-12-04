package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"net/http"
	"time"
)

type TasksHandler struct {
	Service *services.TasksService
}

func NewTasksHandler(service *services.TasksService) *TasksHandler {
	return &TasksHandler{Service: service}
}

// 오늘의 할일 조회 핸들러
func (h *TasksHandler) GetTodayTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetTodayTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// 특정 날짜의 할일 조회 핸들러
func (h *TasksHandler) GetTasksForDate(w http.ResponseWriter, r *http.Request) {
	// 예를 들어, url 파라미터로 날짜를 받는다고 가정
	dateParam := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	tasks, err := h.Service.GetTasksForDate(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// 할일 추가
func (h *TasksHandler) SaveTasks(w http.ResponseWriter, r *http.Request) {
	var tasks models.Tasks

    err := json.NewDecoder(r.Body).Decode(&tasks)
    if err!= nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

	if tasks.Task == "" {
		http.Error(w, "Task cannot be empty", http.StatusBadRequest)
		return
	}

    err = h.Service.SaveTasks(tasks)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}