package models

import "github.com/golang-jwt/jwt/v5"

// Claims defines the JWT claims structure
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
