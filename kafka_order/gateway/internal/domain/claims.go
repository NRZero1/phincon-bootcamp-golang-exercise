package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"Email"`
	Username string `json:"Name"`
	jwt.RegisteredClaims
}