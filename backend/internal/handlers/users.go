package handlers

import (
	"backend/internal/common"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type UsersHandler struct {
	Service *services.UsersService
}

func NewUsersHandler(service *services.UsersService) *UsersHandler {
    return &UsersHandler{Service: service}
}

// 로그인
func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "1002", "Invalid request method", http.StatusMethodNotAllowed)
		return
	} 

	var req struct {
		Email    string `json:"email"`
        Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, "1003", "Invalid input", http.StatusBadRequest)
		return
	}
	

	user, err := h.Service.Login(req.Email, req.Password) 
	if err != nil {
		if errors.Is(err, common.ErrInvalidEmailOrPassword) { 
			utils.ErrorResponse(w, "1001", "Invalid email or password", http.StatusUnauthorized)
		} else { 
			utils.ErrorResponse(w, "5000", "Internal server error", http.StatusInternalServerError)
		}
		return
	}

    utils.SuccessResponse(w, user)
	
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