package auth_test

import (
	"testing"
	"time"

	"github.com/aywa/goNotify/auth"
)

var tokenHashed string = auth.GetToken("test@gmail.com", "test", time.Hour)
var fakeTokenHashed string = "dadadaIAMFAAAKKKEE"

func TestValidateToken(t *testing.T) {
	tokenRead, Claims := auth.ValidateToken(tokenHashed)
	if !tokenRead {
		t.Error("should expect to be true")
	}
	if !(Claims.Email == "test@gmail.com") {
		t.Error("email should be test@gmail.com")
	}
	if !(Claims.UserName == "test") {
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
