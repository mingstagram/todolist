package handlers

// 핸들러는 요청을 처리하고, 서비스 레이어로 연결

import (
	"backend/internal/services"
	"encoding/json"
	"net/http"
)

type TestboardHandler struct {
	service *services.TestboardService
}

func NewTestboardHandler(service *services.TestboardService) *TestboardHandler {
	return &TestboardHandler{service: service}
}

func (h *TestboardHandler) GetAllTestboards(w http.ResponseWriter, r *http.Request) {
	testboards, err := h.service.GetAllTestboards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testboards) 
}