package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID    string  `json:"user_id"`
	TenantID  *string `json:"tenant_id,omitempty"`
	IsInternal bool   `json:"is_internal"`
	jwt.RegisteredClaims
}

func Generate(secret string, claims Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}