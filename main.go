package main

import (
	// "encoding/json"
	"net/http"
	"github.com/mandibchaulagain/food-ecommerce-api-sec-a/responses"
	"github.com/mandibchaulagain/food-ecommerce-api-sec-a/routes"

	
)

func main(){
	print("hello all")
	router := routes.InitAllRoutes()	//creating server
	print("listening on port 8080")
	err :=http.ListenAndServe(":8080",router)
	if(err != nil){
		print("error creating server: ", err.Error())
	}
}

func HandleInitialRoute(w http.ResponseWriter, r *http.Request){
	print("Welcome to go api")
	data := map[string]interface{}{
		"user-list": [] string {"user1","user2","user3"},
		"total-size":100,
		"page":1,
	}
	responses.SuccessResponse(w,data,"Successfully connected")
}
func HandleErrorRoute(w http.ResponseWriter, r *http.Request){
	print("welcome to error route")
	
	responses.ErrorResponse(w,http.StatusInternalServerError,"error seen")
}