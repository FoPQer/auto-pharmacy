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
	r.HandleFunc("/medicines/{medicine}/associate/{tag}", controllers.MedicineAssociateTag).Methods("PUT")

	r.HandleFunc("/supplies", controllers.SupplyIndex).Methods("GET")
	r.HandleFunc("/supplies", controllers.SupplySet).Methods("POST")
	r.HandleFunc("/supplies/{supply}", controllers.SupplyGet).Methods("GET")
	r.HandleFunc("/supplies/{supply}", controllers.SupplyUpdate).Methods("PUT")
	r.HandleFunc("/supplies/{supply}", controllers.SupplyDelete).Methods("DELETE")
	r.HandleFunc("/supplies/{supply}/restore", controllers.SupplyRestore).Methods("PUT")
	r.HandleFunc("/supplies/{supply}/force", controllers.SupplyForceDelete).Methods("DELETE")

	r.HandleFunc("/tags", controllers.TagIndex).Methods("GET")
	r.HandleFunc("/tags", controllers.TagSet).Methods("POST")
	r.HandleFunc("/tags/{tag}", controllers.TagGet).Methods("GET")
	r.HandleFunc("/tags/{tag}", controllers.TagUpdate).Methods("PUT")
	r.HandleFunc("/tags/{tag}", controllers.TagDelete).Methods("DELETE")
	r.HandleFunc("/tags/{tag}/restore", controllers.TagRestore).Methods("PUT")
	r.HandleFunc("/tags/{tag}/force", controllers.TagForceDelete).Methods("DELETE")
	// r.HandleFunc("/tags/{tag}/associate/{medicine}", controllers.MedicineUpdate).Methods("PUT")
}
