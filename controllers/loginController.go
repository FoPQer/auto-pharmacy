package controllers

import (
	"auto-pharmacy/database"
	"auto-pharmacy/models"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()
	var dbUser models.User
	err := database.MysqlDB.DB.Find(&dbUser, "email = ?", email).Error
	if (!ok && errors.Is(err, gorm.ErrRecordNotFound)) || dbUser.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Email or Password doesn't match"))
		return
	}
	if err != nil {
		http.Error(w, "User Find error "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = dbUser.VerifyPassword(password)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Email or Password doesn't match"))
		return
	}
	if err != nil {
		http.Error(w, "Password verify error "+err.Error(), http.StatusInternalServerError)
		return
	}
	tokenLifetime, err := strconv.ParseInt(os.Getenv("TOKEN_LIFETIME"), 10, 32)
	expirationTime := time.Now().Add(time.Duration(tokenLifetime) * time.Minute)
	claims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "auto-pharmacy",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the token to the client
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
