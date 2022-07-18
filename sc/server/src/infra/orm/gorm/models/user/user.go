package user

import (
	"errors"
	"src/infra/orm/gorm/models/base"
	"src/infra/orm/gorm/models/chat"
	"src/infra/orm/gorm/models/message"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	base.Base
	Name     string             `json:"name"`
	Email    string             `gorm:"index:unique"`
	Password string             `json:"_"`
	Messages []*message.Message `gorm:"foreignKey:AuthorId"`
	Chats    []*chat.Chat       `gorm:"many2many:chat_members"`
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
