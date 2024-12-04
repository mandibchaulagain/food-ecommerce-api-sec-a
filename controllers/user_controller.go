package controllers

import (
	"net/http"

	"github.com/mandibchaulagain/food-ecommerce-api-sec-a/modals"
	"github.com/mandibchaulagain/food-ecommerce-api-sec-a/responses"
)
func HandleGetAllUsers(w http.ResponseWriter, r *http.Request){
	allModels := modals.GetAllUsers()
	responses.SuccessResponse(w, allModels, "Successfully fetched user data")
}