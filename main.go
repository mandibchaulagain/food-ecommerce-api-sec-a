package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	print("hello")
	router := mux.NewRouter()//we will use this router only for this whole project
	router.HandleFunc("/api", HandleInitialRoute).Methods("GET")
	//creating server
	print("listening on port 8080")
	err:=http.ListenAndServe(":8080", router)
	if(err!=nil){
		print("error creating server:",err.Error())
	}
}
func HandleInitialRoute(w http.ResponseWriter, r*http.Request){
	print("Welcome to go api by Mandib.")
	w.Header().Set("Content-Type","application/json")//static
	w.WriteHeader(http.StatusOK)//not static, can throw error sometimes
	//sending data...remember this
	data := map[string]interface{}{
		"message":"Welcome to go api laaaa",
		"data":"restart gaeko",
	}
	//_kinda deko re?
	dataByte, _ := json.Marshal(data)
	w.Write(dataByte)//it can be diff sometimes too for response
}