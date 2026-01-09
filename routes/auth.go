package routes

import (
	"auto-pharmacy/controllers"
	"auto-pharmacy/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.Use(middleware.CORSHeadersMiddleware)
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST", "OPTIONS")
}
