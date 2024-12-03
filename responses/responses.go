package responses

import (
	"encoding/json"
	"net/http"
)

func generalResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	dataByte, _ := json.Marshal(data)
	w.Write(dataByte)
}

func SuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	dataWithMessage := map[string]interface{}{
		"data": data,
		"message": message,
	}
	generalResponse(w, http.StatusOK, dataWithMessage)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string){
	dataWithMessage := map[string]interface{}{
		"message": message,
		"data": nil,
	}
	generalResponse(w,statusCode, dataWithMessage)
}