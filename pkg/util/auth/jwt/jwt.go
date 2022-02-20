package jwt

import (
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
	parser := new(jwt.Parser)
	jwtToken, err := parser.Parse(jwtToken, keyFunc)
	if err != nil {
		return "", err
	}
	if claims, ok := jwtToken.Claims.(jwt.StandardClaims); ok && jwtToken.Valid {
		user := claims.Issuer
		return user, nil
	} else {
		return "", errors.New("invalid token")
	}
}
