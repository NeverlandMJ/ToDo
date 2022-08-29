package auth

import "github.com/dgrijalva/jwt-go"

// Create the JWT key used to create the signature
var JwtKey = []byte("my_secret_key")


// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	ID string `json:"id"`
	UserName string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	IsBlocked bool `json:"is_blocked"`
	jwt.StandardClaims
}