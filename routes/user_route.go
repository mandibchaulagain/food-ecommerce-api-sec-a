package routes

import (
	"github.com/gorilla/mux"
	"github.com/mandibchaulagain/food-ecommerce-api-sec-a/controllers"
)

func InitAllUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.HandleGetAllUsers).Methods("GET")
}