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

	// 날짜 포맷 검증 및 변환 (YYYY-MM-DD)
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

// 할일 체크 유무 카운팅
func (h *TasksHandler) CountTasks(w http.ResponseWriter, r *http.Request) { 
	dateParam := r.URL.Query().Get("date")
	// isCheckedParam := r.URL.Query().Get("isChecked")

	// 날짜 포맷 검증 및 변환 (YYYY-MM-DD)
	date, err := time.Parse("2006-01-02", dateParam) 
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	 // isChecked 값 검증 및 변환
	//  var isChecked bool
	//  if isCheckedParam == "true" {
	// 	 isChecked = true
	//  } else if isCheckedParam == "false" {
	// 	 isChecked = false
	//  } else {
	// 	 http.Error(w, "Invalid isChecked value. Must be 'true' or 'false'.", http.StatusBadRequest)
	// 	 return
	//  }

	count, err := h.Service.CountTasks(date)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// 응답 작성
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]int{"count": count})
}