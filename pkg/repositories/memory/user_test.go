package memory

import (
	"testing"

	"example/richard/sovtech/pkg/models"
)

func TestUserRepoCreate(t *testing.T) {
	testUserRepo := NewUserRepository()

	//to force last failure
	firstUser := &models.User{Email: "a@b.c", Password: "1234"}
	firstErr := testUserRepo.Create(firstUser)
	if firstErr != nil {
		t.Fatalf(`expected creation success for user %+v`, firstUser)
	}

	testUsers := []*models.User{
		nil,
		&models.User{},
		&models.User{Email: "a"},
		&models.User{Password: "b"},
		&models.User{Email: "a@b.c", Password: "12345"},
	}

	for _, user := range testUsers {
		err := testUserRepo.Create(user)
		if err == nil {
			t.Fatalf(`expected creation failure for user %+v`, user)
		}
	}
}

func TestGetUser(t *testing.T) {
	testUserRepo := NewUserRepository()
	firstUser := models.NewUser("c@b.c", "1234")
	firstErr := testUserRepo.Create(firstUser)
	if firstErr != nil {
		t.Fatalf(`expected creation success for user %+v`, firstUser)
	}
	secondUser, secondErr := testUserRepo.GetUser("c@b.c")
	if secondErr != nil {
		t.Fatalf(`expected retrieval success, go %s`, secondErr)
	}
	if firstUser.Email != secondUser.Email {
		t.Fatalf(`not the same user: %s != %s`, firstUser.Email, secondUser.Email)
	}
	if firstUser.Password == secondUser.Password {
		t.Fatalf(`password was not hashed: %s == %s`, firstUser.Password, secondUser.Password)
	}
	firstUser.Email = "newemail@test.com"
	if firstUser.Email == secondUser.Email {
		t.Fatalf(`user not immutable: %s == %s`, firstUser.Email, secondUser.Email)
	}
}
