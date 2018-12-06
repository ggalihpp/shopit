package auth

import jwt "github.com/dgrijalva/jwt-go"

// JwtCustomClaim To parse the jwt
type JwtCustomClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
