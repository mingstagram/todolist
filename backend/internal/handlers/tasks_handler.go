package handlers

import (
	_ "backend/internal/kafka"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type UpdateCheckedRequest struct {
	ID int `json:"id"`
	Checked bool `json:"checked"`
}

type TasksHandler struct {
	Service *services.TasksService
}

func NewTasksHandler(service *services.TasksService) *TasksHandler {
	return &TasksHandler{Service: service}
}

// 특정 날짜의 할일 조회 핸들러
func (h *TasksHandler) GetTasksForDate(w http.ResponseWriter, r *http.Request) {
    // 요청 메서드 확인
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// 예를 들어, url 파라미터로 날짜를 받는다고 가정
	dateParam := r.URL.Query().Get("date")
	userId := r.URL.Query().Get("userId")

	// 날짜 포맷 검증 및 변환 (YYYY-MM-DD)
	date, err := time.Parse("2006-01-02", dateParam) 
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	tasks, err := h.Service.GetTasksForDate(date, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// 할일 추가
func (h *TasksHandler) SaveTasks(w http.ResponseWriter, r *http.Request) {
    // 요청 메서드 확인
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "1002", "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var tasks models.Tasks 

    err := json.NewDecoder(r.Body).Decode(&tasks) 
    if err!= nil {
        utils.ErrorResponse(w, "1003", "Invalid input", http.StatusBadRequest)
        return
    }
	if tasks.Task == "" {
		utils.ErrorResponse(w, "1004", "Task cannot be empty", http.StatusBadRequest)
		return
	}
    err = h.Service.SaveTasks(tasks)
    if err!= nil {
        utils.ErrorResponse(w, "5000", err.Error(), http.StatusInternalServerError)
        return
    }

    utils.CreatedResponse(w, []models.Tasks{})
}

// 할일 체크 유무 카운팅
func (h *TasksHandler) CountTasks(w http.ResponseWriter, r *http.Request) {  
    // 요청 메서드 확인
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, "1002",  "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
 
	dateParam := r.URL.Query().Get("date")
	userId := r.URL.Query().Get("userId")
	// isCheckedParam := r.URL.Query().Get("isChecked")

	// 날짜 포맷 검증 및 변환 (YYYY-MM-DD)
	date, err := time.Parse("2006-01-02", dateParam) 
	if err != nil {
		utils.ErrorResponse(w, "1005",  "Invalid date format", http.StatusBadRequest)
		return
	}  

	count, err := h.Service.CountTasks(date, userId)
	if err != nil {
        utils.ErrorResponse(w, "5000",  err.Error(), http.StatusInternalServerError)
        return
    } 
 
	utils.SuccessResponse(w, map[string]int{"count": count}) 
}

// 할일 체크 / 체크해제
func (h *TasksHandler) UpdateChecked(w http.ResponseWriter, r *http.Request) {
    // 요청 메서드 확인
	if r.Method != http.MethodPut {
		utils.ErrorResponse(w, "1002", "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// JSON 요청 본문 파싱
	var req UpdateCheckedRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "1006", "Invalid request body", http.StatusBadRequest)
		return
	}

	// 요청 데이터 사용
	err = h.Service.UpdateChecked(req.Checked, req.ID)
	if err != nil {
		utils.ErrorResponse(w, "1007", fmt.Sprintf("Failed to update task: %v", err), http.StatusInternalServerError)
		return
	}

	// // Kafka로 상태 변경 메시지 전송
	// message := fmt.Sprintf("Task ID: %d updated to checked %v", req.ID, req.Checked)
	// err = kafka.SendMessage(message)
	// if err != nil {
	// 	utils.ErrorResponse(w, "5000", "Failed to send kafka message", http.StatusInternalServerError)
	// 	return
	// } 

	utils.SuccessResponse(w, "Task updated successfully")  
}

// 할일 삭제
func (h *TasksHandler) DeleteTasks(w http.ResponseWriter, r *http.Request) { 
	if r.Method != http.MethodDelete {
		utils.ErrorResponse(w, "1002", "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		utils.ErrorResponse(w, "1008", "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ErrorResponse(w, "1008", "Invalid id parameter", http.StatusBadRequest)
		return
	}

	// 할일 삭제
	err = h.Service.DeleteTasks(id)
	if err != nil {
		utils.ErrorResponse(w, "5000", "Failed to delete tasks", http.StatusInternalServerError)
		return
	}

	// // Kafka로 삭제된 할일 메시지 전송
	// message := fmt.Sprintf("Task ID: %d deleted", id)
	// err = kafka.SendMessage(message)
	// if err != nil {
	// 	utils.ErrorResponse(w, "5000", "Failed to send kafka message", http.StatusInternalServerError)
    //     return
	// }

	utils.NoContentResponse(w)
}