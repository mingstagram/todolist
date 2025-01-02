package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
    return &AuthHandler{Service: service}
}

var Req struct {
	RedirectURI string `json:"redirect_uri"`
	ClientID string `json:"client_id"`
	ResponseType string `json:"response_type"` 
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
        Password string `json:"password"`
	}
	
	fmt.Println("Received POST request for /auth/signup")
	
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

// 카카오 로그인
func (h *AuthHandler) InitiateKakaoLogin(w http.ResponseWriter, r *http.Request) { 
	Req := struct {
		RedirectURI  string `json:"redirect_uri"`
		ClientID     string `json:"client_id"`
		ResponseType string `json:"response_type"`
	}{
		RedirectURI:  "http://localhost:3000/kakao/callback",
		ClientID:     "f935795848dcfc4ebec4716acfc63bfa",
		ResponseType: "code",
	}

	// 요청이 GET이 아닐 경우 에러 반환
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, "1011", "Invalid kakao url", http.StatusBadRequest)
		return
	}

	// 클라이언트에 JSON 응답
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Req)
}

func (h *AuthHandler) HandleKakaoCallback(w http.ResponseWriter, r *http.Request) {  
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, "1011", "Invalid kakao url", http.StatusBadRequest)
		return
	}

	// URL에서 code 추출
	code := r.URL.Query().Get("code")

	// 카카오 서비스 호출하여 access token 얻기
	accessToken, err := h.Service.GetKakaoAccessToken(code)
	if err != nil {
		utils.ErrorResponse(w, "1012", "Failed to get access token", http.StatusInternalServerError)
		return
	} 

	_, err = h.Service.GetKakaoUserInfo(accessToken)

	// 여기서 추가로, accessToken을 사용하여 사용자 정보를 가져오거나 회원가입을 진행할 수 있습니다.

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]string{"access_token": accessToken})
}