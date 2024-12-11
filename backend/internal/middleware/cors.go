package middleware

import (
	"net/http"
)

// CORS 미들웨어 함수
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS 헤더 설정
		w.Header().Set("Access-Control-Allow-Origin", "*") // 모든 Origin을 허용
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type") // 허용할 헤더
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 허용할 HTTP 메서드 설정

		// 만약 OPTIONS 요청이면 바로 응답
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 실제 요청 처리
		next.ServeHTTP(w, r)
	})
}
