package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
    return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
        Password string `json:"password"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ErrorResponse(w, "1003", "Invalid input", http.StatusBadRequest)
        return
	} 
	responseData, err := h.Service.Login(payload.Email, payload.Password)
	if err!= nil {
        utils.ErrorResponse(w, "1005", "Invalid credentials", http.StatusUnauthorized)
        return
    }

	utils.SuccessResponse(w, responseData) 
}

// 회원가입
func (h *AuthHandler) SaveUsers(w http.ResponseWriter, r *http.Request) {
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