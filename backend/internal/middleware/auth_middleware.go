package middleware

import (
	"backend/internal/utils" 
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 요청 헤더 전체 출력 
		authHeader := r.Header.Get("Authorization")  
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.ErrorResponse(w, "1010", "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			utils.ErrorResponse(w, "1011", "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}