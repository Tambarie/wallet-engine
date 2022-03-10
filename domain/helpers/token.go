package helpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(method *jwt.SigningMethodHMAC, claim jwt.MapClaims, secret string) (string, error) {
	return jwt.NewWithClaims(method, claim).SignedString([]byte(secret))
}

func GenerateClaims(email, ID string) (jwt.MapClaims, jwt.MapClaims) {
	accessClaim := jwt.MapClaims{
		"email": email,
		"id":    ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	refreshClaim := jwt.MapClaims{
		"id":  ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"sub": 1,
	}

	return accessClaim, refreshClaim
}

func VerifyToken(token string, claims jwt.MapClaims, secret string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
