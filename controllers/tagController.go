package controllers

import (
	"auto-pharmacy/services"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TagIndex(w http.ResponseWriter, r *http.Request) {
	tags, err := services.GetAllTags()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(res); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func TagGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag, err := services.GetTag(vars["tag"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(tag)
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

func TagSet(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	tag, err := services.SetTag(body)
	if err != nil {
		http.Error(w, "Tag error "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(tag)
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

func TagUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := json.NewDecoder(r.Body)
	tag, err := services.UpdateMedicine(vars["tag"], body)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Tag error "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(tag)
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

func TagDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.DeleteMedicine(vars["tag"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Tag error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func TagRestore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.RestoreMedicine(vars["tag"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Tag error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func TagForceDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.ForceDeleteMedicine(vars["tag"])
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Tag error "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write(nil); err != nil {
		http.Error(w, "Send response error "+err.Error(), http.StatusInternalServerError)
		return
	}
}
