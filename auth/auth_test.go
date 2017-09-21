package auth_test

import (
	"testing"
	"time"

	"github.com/aywa/goNotify/auth"
)

var tokenHashed, err = auth.GetToken("test@gmail.com", "test", time.Hour)
var fakeTokenHashed string = "dadadaIAMFAAAKKKEE"

func TestValidateToken(t *testing.T) {
	tokenRead, Claims := auth.ValidateToken(tokenHashed)
	if !tokenRead {
		t.Error("should expect to be true")
	}
	if !(Claims.Email == "test@gmail.com") {
		t.Error("email should be test@gmail.com")
	}
	if !(Claims.FirstName == "test") {
		t.Error("userName should be test")
	}

	fakeTokenRead, fakeClaims := auth.ValidateToken(fakeTokenHashed)
	if fakeTokenRead {
		t.Error("should expect to be false")
	}
	if !(fakeClaims == nil) {
		t.Error("should expect to be nil")
	}
	println(fakeClaims)
}

var password = "myTest13213Password."

func TestSaltPasswordAndCompare(t *testing.T) {
	saltP, err := auth.SaltPassword(password)
	if err != nil {
		t.Error("should not have error when salt a string")
	}
	err = auth.CompareHashAndPassword([]byte(saltP), []byte(password))
	if err != nil {
		t.Error("should not return an error if we compare the password and the salt one")
	}
}

func TestCompareWrongPassword(t *testing.T) {
	fakeHP, _ := auth.SaltPassword("random")
	err = auth.CompareHashAndPassword([]byte(fakeHP), []byte(password))
	if err == nil {
		t.Error("should return an error if we compare wrong salt and a password")
	}
}

func TestGetUserFromAuth(t *testing.T) {
	c, err := auth.GetUserFromAuth(tokenHashed)
	if err != nil {
		t.Error("should not return an error if token is valid")
	}
	if c.Email != "test@gmail.com" && c.FirstName != "test" {
		t.Error("should return a Claims with the good information")
	}
	_, err = auth.GetUserFromAuth(fakeTokenHashed)

	if err == nil {
		t.Error("should return an error if token is not valid")
	}
}
