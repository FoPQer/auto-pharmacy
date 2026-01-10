package controllers

import (
	"auto-pharmacy/services"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SupplyIndex(w http.ResponseWriter, r *http.Request) {
	supplies, err := services.GetAllSupplies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(supplies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(res); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func SupplyGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supply, err := services.GetSupply(vars["supply"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(supply)
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

func SupplySet(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	sup, err := services.SetSupply(body)
	if err != nil {
		http.Error(w, "Supply error "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(sup)
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

func SupplyUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := json.NewDecoder(r.Body)
	sup, err := services.UpdateSupply(vars["supply"], body)
	if err != nil {
		http.Error(w, "Supply error "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(sup)
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

func SupplyDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := services.DeleteSupply(vars["supply"]); err != nil {
		http.Error(w, "Supply error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func SupplyRestore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := services.RestoreSupply(vars["supply"]); err != nil {
		http.Error(w, "Supply error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func SupplyForceDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := services.ForceDeleteSupply(vars["supply"]); err != nil {
		http.Error(w, "Supply error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func SupplyMassDelete(w http.ResponseWriter, r *http.Request) {
	var data map[string][]int
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Body parse error "+err.Error(), http.StatusInternalServerError)
		return
	}
	ids, ok := data["ids"]
	if !ok {
		http.Error(w, "Data find error ", http.StatusInternalServerError)
		return
	}
	err := services.MassDeleteSupply(ids)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Supply error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}
