package routes

import "github.com/gorilla/mux"

func InitAllRoutes() *mux.Router {
	newRouter := mux.NewRouter()
	InitAllPracticeRoutes(newRouter)
	InitAllUserRoutes(newRouter)
	return newRouter
}