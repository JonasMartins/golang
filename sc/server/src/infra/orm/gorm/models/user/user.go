package user

import (
	"errors"
	base "src/infra/orm/gorm/models/base"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	base.Base
	Name     string `json:"name"`
	Email    string `gorm:"index:unique"`
	Password string `json:"_"`
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
