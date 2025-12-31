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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(medicines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func MedicineGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	medicine, err := services.GetMedicine(vars["medicine"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(medicine)
	if errors.Is(err, errors.New("record not found")) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func MedicineSet(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	med, err := services.SetMedicine(body)
	if err != nil {
		http.Error(w, "Medicine error "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(med)
	if err != nil {
		http.Error(w, "Marshall error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(res); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func MedicineUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := json.NewDecoder(r.Body)
	med, err := services.UpdateMedicine(vars["medicine"], body)
	if err != nil {
		http.Error(w, "Medicine error "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(med)
	if err != nil {
		http.Error(w, "Marshall error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func MedicineDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := services.DeleteMedicine(vars["medicine"]); err != nil {
		http.Error(w, "Medicine error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func MedicineRestore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := services.RestoreMedicine(vars["medicine"]); err != nil {
		http.Error(w, "Medicine error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func MedicineForceDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := services.ForceDeleteMedicine(vars["medicine"]); err != nil {
		http.Error(w, "Medicine error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}
