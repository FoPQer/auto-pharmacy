package controllers

import (
	"auto-pharmacy/services"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func MedicineIndex(w http.ResponseWriter, r *http.Request) {
	medicines, err := services.GetAllMedicines()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	res, err := json.Marshal(medicines)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.Write(res)
}

func MedicineGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	medicine, err := services.GetMedicine(vars["medicine"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	res, err := json.Marshal(medicine)
	if errors.Is(err, errors.New("record not found")) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.Write(res)
}
