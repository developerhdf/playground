package auth

import (
	"context"
	"net/http"

	"example/richard/sovtech/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

var ctxUserKey = "user"

func PasswordHash(password string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwdHash), err
}

func ValidPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AddUserToRequestContext(r *http.Request, user *models.User) {
	if r == nil {
		return
	}
	ctx := context.WithValue(r.Context(), ctxUserKey, user)
	r = r.WithContext(ctx)
}

func GetUserFromRequestContext(r *http.Request) *models.User {
	var user *models.User
	if r != nil {
		ctx := r.Context()
		if ctxUser := ctx.Value(ctxUserKey); ctxUser != nil {
			if modelUser, isModelUser := ctxUser.(*models.User); isModelUser {
				user = modelUser
			}
		}
	}
	return user
}
