package responses

import ("net/http" 
"encoding/json")

//private when small letter starting
// data type where rakhne?
//generalResponse returns statusCode and data
func generalResponse(w http.ResponseWriter, statusCode int, data interface{}){
	w.Header().Set("Content-Type","application/json")//tells api consumer response will be in json format
	w.WriteHeader((statusCode))//sends SC to client if request was successful or not
	dataByte, _:= json.Marshal(data)//converts data of any type into json format, dataByte is a json-encoded data
	//the _ ignores the error returned by json.marshal
	w.Write(dataByte)//writes dataByte to response body
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