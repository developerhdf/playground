package auth

import (
	"context"
	"net/http"

	"example/richard/sovtech/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func PasswordHash(password string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwdHash), err
}

func ValidPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AddUserToRequestContext(r *http.Request, user *models.User) *http.Request {
	if r == nil || user == nil {
		return r
	}
	ctx := context.WithValue(r.Context(), userCtxKey, user)
	return r.WithContext(ctx)
}

func GetUserFromContext(ctx context.Context) *models.User {
	var user *models.User
	if ctxUser := ctx.Value(userCtxKey); ctxUser != nil {
		if modelUser, isModelUser := ctxUser.(*models.User); isModelUser {
			user = modelUser
		}
	}
	return user
}
