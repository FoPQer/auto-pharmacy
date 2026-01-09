package controllers

import (
	"auto-pharmacy/services"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := services.GetUser(vars["user"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(user)
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

func UserSet(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	med, err := services.SetUser(body)
	if err != nil {
		http.Error(w, "User error "+err.Error(), http.StatusInternalServerError)
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

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := json.NewDecoder(r.Body)
	med, err := services.UpdateUser(vars["user"], body)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "User error "+err.Error(), http.StatusInternalServerError)
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

func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.DeleteUser(vars["user"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "User error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserRestore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.RestoreUser(vars["user"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "User error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserForceDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.ForceDeleteUser(vars["user"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "User error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserMassDelete(w http.ResponseWriter, r *http.Request) {
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
	err := services.MassDeleteUser(ids)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "User error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}
