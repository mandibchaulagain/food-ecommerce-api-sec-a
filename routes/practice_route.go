package routes

import (
	"github.com/gorilla/mux"
	"github.com/mandibchaulagain/food-ecommerce-api-sec-a/controllers"
)

func InitAllPracticeRoutes(router *mux.Router){
	router.HandleFunc("/error",controllers.HandleErrorRequest)
	router.HandleFunc("/success",controllers.HandleSuccessRequest)
}