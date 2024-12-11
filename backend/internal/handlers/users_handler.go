package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

type UsersHandler struct {
	Service *services.UsersService
}

func NewUsersHandler(service *services.UsersService) *UsersHandler {
    return &UsersHandler{Service: service}
}  

// 회원가입
func (h *UsersHandler) SaveUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "1002", "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.Users 
	err := json.NewDecoder(r.Body).Decode(&user) 
    if err!= nil {
        utils.ErrorResponse(w, "1003", "Invalid input", http.StatusBadRequest)
        return
    }

	if err = h.Service.SaveUsers(user); err != nil {
		utils.ErrorResponse(w, "5001", "Failed to save user", http.StatusInternalServerError)
		return
	} 

    utils.SuccessResponse(w, "User registered successfully")
}