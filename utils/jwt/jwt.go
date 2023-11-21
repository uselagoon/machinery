package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// LagoonClaims is a set of JWT claims used by Lagoon.
type LagoonClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

// GenerateAdminToken returns a JWT admin token valid for the number of seconds requested
func GenerateAdminToken(secret, audience, subject, issuer string, issued, expireSeconds int64) (string, error) {
	if issued == 0 {
		issued = time.Now().Unix()
	}
	claims := LagoonClaims{
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			Audience:  audience,
			ExpiresAt: issued + expireSeconds,
			IssuedAt:  issued,
			Subject:   subject,
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
