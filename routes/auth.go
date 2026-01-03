package routes

import (
	"auto-pharmacy/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
}
