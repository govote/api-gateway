package models

import "github.com/dgrijalva/jwt-go"

// Claims Jwt token required claims
type Claims struct {
	jwt.StandardClaims
	Name       string `json:"name"`
	FacebookID string `json:"facebookId"`
}
