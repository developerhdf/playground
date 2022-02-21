package graph

import (
	"context"
	"fmt"

	"example/richard/sovtech/pkg/models"
	"example/richard/sovtech/pkg/repositories"
	"example/richard/sovtech/pkg/util/auth"
	"example/richard/sovtech/pkg/util/auth/jwt"

	"example/richard/sovtech/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PeopleRepo repositories.PeopleRepository
	UserRepo   repositories.UserRepository
}

func NewSovTechResolver(ur repositories.UserRepository, pr repositories.PeopleRepository) *Resolver {
	return &Resolver{
		PeopleRepo: pr,
		UserRepo:   ur,
	}
}

func GetContextUser(ctx context.Context) *models.User {
	return auth.GetUserFromContext(ctx)
}

func ToModelsUser(user model.NewUser) *models.User {
	return models.NewUser(user.Email, user.Password)
}

func GetJWTToken(user *models.User) (string, error) {
	if user == nil {
		return "", fmt.Errorf("user is nil")
	}
	return jwt.CreateToken(user.Email)
}

func CreateToken(username string) (string, error) {
	return jwt.CreateToken(username)
}

func GetTokenUsername(token string) (string, error) {
	return jwt.GetTokenUsername(token)
}

func ValidPassword(password, hash string) bool {
	return auth.ValidPassword(password, hash) == nil
}
