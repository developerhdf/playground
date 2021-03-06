package memory

import (
	"errors"
	"sync"

	"example/richard/sovtech/pkg/models"
	"example/richard/sovtech/pkg/util/auth"
)

const (
	MaxUsers = 20
)

type UserRepository struct {
	userStore map[string]*models.User
	createMu  sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userStore: make(map[string]*models.User),
	}
}

func (ur *UserRepository) Create(user *models.User) error {
	var err error
	ur.createMu.Lock()
	defer ur.createMu.Unlock()
	switch {
	case user == nil || user.Email == "" || user.Password == "":
		err = errors.New("invalid user")
	case len(ur.userStore) < MaxUsers:
		if _, found := ur.userStore[user.Email]; !found {
			if passwordHash, err := auth.PasswordHash(user.Password); err == nil {
				userCopy := *user
				userCopy.Password = passwordHash
				ur.userStore[user.Email] = &userCopy
			}
		} else {
			err = errors.New("user exists")
		}
	default:
		err = errors.New("max supported users reached")
	}
	return err
}

func (ur *UserRepository) GetUser(username string) (*models.User, error) {
	var user *models.User
	var err error
	if memUser, ok := ur.userStore[username]; ok && memUser != nil && memUser.Active {
		user = memUser
	} else {
		err = errors.New("user not found")
	}
	return user, err
}

func (ur *UserRepository) GetPasswordHash(username string) (string, error) {
	var passwordHash string
	var err error
	if memUser, ok := ur.userStore[username]; ok && memUser != nil && memUser.Active {
		passwordHash = memUser.Password
	} else {
		err = errors.New("user not found")
	}
	return passwordHash, err
}
