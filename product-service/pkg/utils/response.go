package utils

import (
	"encoding/json"
	"net/http"

	"github.com/mangesh-shinde/ecomm-app/product-service/pkg/models"
)

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := models.ErrorResponse{ErrorMessage: err.Error()}
	json.NewEncoder(w).Encode(errorResponse)
}
