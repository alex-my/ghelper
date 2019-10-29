package jwt_test

import (
	"testing"

	"github.com/alex-my/ghelper/jwt"
)

func TestIPToInt(t *testing.T) {
	key := []byte("123456")

	token, err := jwt.TokenWithKey(key)
	if err != nil {
		t.Fatalf("create token failed: %s", err.Error())
	}
	t.Logf("token: %s", token)

	p, err := jwt.Verify(token)
	if err != nil {
		t.Fatalf("verify token failed: %s", err.Error())
	}

	t.Logf("exp: %d, iat: %d", p["exp"], p["iat"])
}
