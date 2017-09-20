package auth

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
	BCRYPT_COST   = 1
)

var secret = []byte("mySuperSecretYolo")

// Claims is the struct of or jwt
type Claims struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	// recommended having
	jwt.StandardClaims
}

// GetToken create a token from a email, and name
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

// GetUserFromAuth Return User if the
func GetUserFromAuth(a string) error {
	return errors.New("user is not auth")
}

// SaltPassword use bcrypt to salt a password
func SaltPassword(p string) (string, error) {
	a, err := bcrypt.GenerateFromPassword([]byte(p), BCRYPT_COST)
	return string(a), err
}

// CompareHashAndPassword use bcrypt to compare a password
func CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
