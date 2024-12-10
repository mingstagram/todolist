package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, code, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]interface{} {
		"code": code,
		"message": message,
	}
	json.NewEncoder(w).Encode(response)
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{} {
		"code": "0000",
		"message": "Success",
		"data": data,
	}
	json.NewEncoder(w).Encode(response)
}

func CreatedResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"code":    "0000",
		"message": "Resource created successfully",
		"data":    data,
	}
	json.NewEncoder(w).Encode(response)
}

func NoContentResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) 
}