package models

import (
	"errors"
	"src/infra/orm/gorm/models/base"

	"golang.org/x/crypto/bcrypt"
)

type Chat struct {
	base.Base
	Members  []*User    `gorm:"many2many:chat_members;"`
	Messages []*Message `gorm:"foreigney:ChatId"`
}

type User struct {
	base.Base
	Name     string     `json:"name"`
	Email    string     `gorm:"index:unique"`
	Password string     `json:"_"`
	Messages []*Message `gorm:"foreignKey:AuthorId"`
	Chats    []*Chat    `gorm:"many2many:chat_members"`
}

type Message struct {
	base.Base
	Body     string `json:"body"`
	AuthorId string `json:"authorId"`
	ChatId   string `json:"chatId"`
	Seen     bool   `json:"seen"`
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
