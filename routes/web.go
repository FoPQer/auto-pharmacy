package routes

import (
	"auto-pharmacy/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/medicines", controllers.MedicineIndex).Methods("GET")
	r.HandleFunc("/medicines", controllers.MedicineSet).Methods("POST")
	r.HandleFunc("/medicines/{medicine}", controllers.MedicineGet).Methods("GET")
	r.HandleFunc("/medicines/{medicine}", controllers.MedicineUpdate).Methods("PUT")
	r.HandleFunc("/medicines/{medicine}", controllers.MedicineDelete).Methods("DELETE")
	r.HandleFunc("/medicines/{medicine}/restore", controllers.MedicineRestore).Methods("PUT")
	r.HandleFunc("/medicines/{medicine}/force", controllers.MedicineForceDelete).Methods("DELETE")
	r.HandleFunc("/medicines/{medicine}/release", controllers.MedicineRelease).Methods("GET")

	r.HandleFunc("/supply", controllers.SupplyIndex).Methods("GET")
	r.HandleFunc("/supply", controllers.SupplySet).Methods("POST")
	r.HandleFunc("/supply/{supply}", controllers.SupplyGet).Methods("GET")
	r.HandleFunc("/supply/{supply}", controllers.SupplyUpdate).Methods("PUT")
	r.HandleFunc("/supply/{supply}", controllers.SupplyDelete).Methods("DELETE")
	r.HandleFunc("/supply/{supply}/restore", controllers.SupplyRestore).Methods("PUT")
	r.HandleFunc("/supply/{supply}/force", controllers.SupplyForceDelete).Methods("DELETE")
}
