package responses

import ("net/http" 
"encoding/json")

//private when small letter starting
// data type where
func generalResponse(w http.ResponseWriter, statusCode int, data interface{}){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader((statusCode))
	dataByte, _:= json.Marshal(data)
	w.Write(dataByte)
}

func SuccessResponse(w http.ResponseWriter, data interface{}, message string){
	dataWithMessage:=map[string]interface{}{
		"data":data,
		"message":message,
	}
	generalResponse(w, http.StatusOK, dataWithMessage)
}

func ErrorResponse(){
	
}