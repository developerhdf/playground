package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var Secret = []byte("#3B8x0L1sT5fMq")

func keyFunc(token *jwt.Token) (interface{}, error) {
	return Secret, nil
}

func CreateToken(user string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		Issuer:    user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

func GetTokenUsername(token string) (string, error) {
	user := ""
	var tokenErr error
	parser := new(jwt.Parser)
	if jwtToken, err := parser.ParseWithClaims(token, &jwt.StandardClaims{}, keyFunc); err == nil {
		if claims, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			user = claims.Issuer
		} else {
			tokenErr = errors.New("invalid token")
		}
	} else {
		tokenErr = err
	}
	return user, tokenErr
}
