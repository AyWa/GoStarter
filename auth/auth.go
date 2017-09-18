package goAuth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte("mySuperSecretYolo")

// Claims is the struct of or jwt
type Claims struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	// recommended having
	jwt.StandardClaims
}

// GetToken is token from a email, and name
func GetToken(email, name string, expire time.Duration) string {
	// Expires the token and cookie in expire
	expireToken := time.Now().Add(expire).Unix()

	claims := Claims{
		email,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
		},
	}
	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString(secret)

	return signedToken
}

// ValidateToken a token from a email, and name
func ValidateToken(tokenReceived string) (bool, *Claims) {
	token, err := jwt.ParseWithClaims(tokenReceived, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return secret, nil
	})
	if err != nil {
		return false, nil
	}

	// Grab the tokens claims and pass it into the original request
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return true, claims
	}
	return false, nil
}
