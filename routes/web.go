package routes

import (
	"auto-pharmacy/controllers"

	"github.com/gorilla/mux"
)

func RegisterWebRoutes(r *mux.Router) {
	medicineRoutes(r.PathPrefix("/medicines").Subrouter())
	supplyRoutes(r.PathPrefix("/supplies").Subrouter())
	tagRoutes(r.PathPrefix("/tags").Subrouter())
	userRoutes(r.PathPrefix("/users").Subrouter())
}

func medicineRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.MedicineIndex).Methods("GET")
	r.HandleFunc("", controllers.MedicineSet).Methods("POST")
	r.HandleFunc("/{medicine}", controllers.MedicineGet).Methods("GET")
	r.HandleFunc("/{medicine}", controllers.MedicineUpdate).Methods("PUT")
	r.HandleFunc("/{medicine}", controllers.MedicineDelete).Methods("DELETE")
	r.HandleFunc("/{medicine}/restore", controllers.MedicineRestore).Methods("PUT")
	r.HandleFunc("/{medicine}/force", controllers.MedicineForceDelete).Methods("DELETE")
	r.HandleFunc("/{medicine}/release", controllers.MedicineRelease).Methods("GET")
	r.HandleFunc("/{medicine}/associate/{tag}", controllers.MedicineAssociateTag).Methods("PUT")
}

func supplyRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.SupplyIndex).Methods("GET")
	r.HandleFunc("", controllers.SupplySet).Methods("POST")
	r.HandleFunc("/{supply}", controllers.SupplyGet).Methods("GET")
	r.HandleFunc("/{supply}", controllers.SupplyUpdate).Methods("PUT")
	r.HandleFunc("/{supply}", controllers.SupplyDelete).Methods("DELETE")
	r.HandleFunc("/{supply}/restore", controllers.SupplyRestore).Methods("PUT")
	r.HandleFunc("/{supply}/force", controllers.SupplyForceDelete).Methods("DELETE")
}

func tagRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.TagIndex).Methods("GET")
	r.HandleFunc("", controllers.TagSet).Methods("POST")
	r.HandleFunc("/{tag}", controllers.TagGet).Methods("GET")
	r.HandleFunc("/{tag}", controllers.TagUpdate).Methods("PUT")
	r.HandleFunc("/{tag}", controllers.TagDelete).Methods("DELETE")
	r.HandleFunc("/{tag}/restore", controllers.TagRestore).Methods("PUT")
	r.HandleFunc("/{tag}/force", controllers.TagForceDelete).Methods("DELETE")
	// r.HandleFunc("/{tag}/associate/{medicine}", controllers.MedicineUpdate).Methods("PUT")
}

func userRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.UserIndex).Methods("GET")
	r.HandleFunc("", controllers.UserSet).Methods("POST")
	r.HandleFunc("/{user}", controllers.UserGet).Methods("GET")
	r.HandleFunc("/{user}", controllers.UserUpdate).Methods("PUT")
	r.HandleFunc("/{user}", controllers.UserDelete).Methods("DELETE")
	r.HandleFunc("/{user}/restore", controllers.UserRestore).Methods("PUT")
	r.HandleFunc("/{user}/force", controllers.UserForceDelete).Methods("DELETE")
}
