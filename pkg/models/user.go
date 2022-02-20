package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Active   bool   `json:"-"`
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
		Active:   true,
	}
}
