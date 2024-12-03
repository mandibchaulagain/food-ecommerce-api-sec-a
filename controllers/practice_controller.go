package controllers

import (
	"net/http"

	"github.com/mandibchaulagain/food-ecommerce-api-sec-a/responses"
)

func HandleSuccessRequest(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"users": [] string {"user1","user2","user3"},
		"total-size":100,
		"page":1,
	}
	responses.SuccessResponse(w,data,"Successfully fetched user")
}
func HandleErrorRequest(w http.ResponseWriter, r *http.Request){
	responses.ErrorResponse(w, http.StatusUnauthorized, "Not Authorized to access this route")
}