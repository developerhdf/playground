package repositories

import (
	"example/richard/sovtech/pkg/models"
)

type UserRepository interface {
	Create(user *models.User) (string, error)
	GetUser(username string) (*models.User, error)
	GetPasswordHash(username string) (string, error)
}
