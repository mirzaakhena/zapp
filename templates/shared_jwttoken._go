package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// IToken is
type IToken interface {
	GenerateToken(subject, issuer, audience string, extendData interface{}, expiredInHour uint) string
	ValidateToken(subject, token string) *CustomJwt
}

// JwtToken is
type JwtToken struct {
	SecretKey []byte
}

func NewToken() *JwtToken {
	return &JwtToken{}
}

// CustomJwt is
type CustomJwt struct {
	jwt.StandardClaims
	ExtendData interface{}
}

// GenerateToken is
func (j *JwtToken) GenerateToken(subject, issuer, audience string, extendData interface{}, expiredInHour uint) string {
	t := time.Now().UTC()
	claims := &CustomJwt{
		ExtendData: extendData,
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			Issuer:    issuer,
			Audience:  audience,
			IssuedAt:  t.Unix(),
			ExpiresAt: t.Add(time.Duration(expiredInHour) * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(j.SecretKey)
	return tokenString
}

// ValidateToken is
func (j *JwtToken) ValidateToken(subject, tokenString string) *CustomJwt {

	token, err := jwt.ParseWithClaims(tokenString, &CustomJwt{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.SecretKey, nil
	})

	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(*CustomJwt)
	if !ok || !token.Valid {
		// invalid token
		return nil
	}

	if claims.Subject != subject {
		// wrong subject
		return nil
	}

	return claims
}
