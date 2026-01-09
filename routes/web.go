package routes

import (
	"auto-pharmacy/controllers"
	"auto-pharmacy/middleware"

	"github.com/gorilla/mux"
)

func RegisterWebRoutes(r *mux.Router) {
	r.Use(middleware.CORSHeadersMiddleware)
	r.Use(middleware.JWTAuthMiddleware)
	medicineRoutes(r.PathPrefix("/medicines").Subrouter())
	supplyRoutes(r.PathPrefix("/supplies").Subrouter())
	tagRoutes(r.PathPrefix("/tags").Subrouter())
	userRoutes(r.PathPrefix("/users").Subrouter())
}

func medicineRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.MedicineIndex).Methods("GET", "OPTIONS")
	r.HandleFunc("", controllers.MedicineSet).Methods("POST", "OPTIONS")
	r.HandleFunc("/{medicine}", controllers.MedicineGet).Methods("GET", "OPTIONS")
	r.HandleFunc("/{medicine}", controllers.MedicineUpdate).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{medicine}", controllers.MedicineDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/{medicine}/restore", controllers.MedicineRestore).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{medicine}/force", controllers.MedicineForceDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/{medicine}/release", controllers.MedicineRelease).Methods("GET", "OPTIONS")
	r.HandleFunc("/{medicine}/associate/{tag}", controllers.MedicineAssociateTag).Methods("PUT", "OPTIONS")
}

func supplyRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.SupplyIndex).Methods("GET", "OPTIONS")
	r.HandleFunc("", controllers.SupplySet).Methods("POST", "OPTIONS")
	r.HandleFunc("/{supply}", controllers.SupplyGet).Methods("GET", "OPTIONS")
	r.HandleFunc("/{supply}", controllers.SupplyUpdate).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{supply}", controllers.SupplyDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/{supply}/restore", controllers.SupplyRestore).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{supply}/force", controllers.SupplyForceDelete).Methods("DELETE", "OPTIONS")
}

func tagRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.TagIndex).Methods("GET", "OPTIONS")
	r.HandleFunc("", controllers.TagSet).Methods("POST", "OPTIONS")
	r.HandleFunc("/{tag}", controllers.TagGet).Methods("GET", "OPTIONS")
	r.HandleFunc("/{tag}", controllers.TagUpdate).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{tag}", controllers.TagDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/{tag}/restore", controllers.TagRestore).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{tag}/force", controllers.TagForceDelete).Methods("DELETE", "OPTIONS")
	// r.HandleFunc("/{tag}/associate/{medicine}", controllers.MedicineUpdate).Methods("PUT", "OPTIONS")
}

func userRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.UserIndex).Methods("GET", "OPTIONS")
	r.HandleFunc("", controllers.UserSet).Methods("POST", "OPTIONS")
	r.HandleFunc("/{user}", controllers.UserGet).Methods("GET", "OPTIONS")
	r.HandleFunc("/{user}", controllers.UserUpdate).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{user}", controllers.UserDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/{user}/restore", controllers.UserRestore).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{user}/force", controllers.UserForceDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/delete/mass", controllers.UserMassDelete).Methods("POST", "OPTIONS")
}
