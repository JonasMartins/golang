package models

import (
	"errors"
	"src/infra/orm/gorm/models/base"
	"sync"

	"github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

type MessageObserver struct {
	UserId  string
	Message chan *Message
}

type Chat struct {
	base.Base
	Members          []*User    `gorm:"many2many:chat_members;"`
	Messages         []*Message `gorm:"foreigney:ChatId"`
	MessageObservers sync.Map   `gorm:"-:all"`
}

type User struct {
	base.Base
	Name     string     `json:"name"`
	Email    string     `gorm:"index:unique"`
	Password string     `json:"_"`
	Profile  string     `json:"profile,omitempty"`
	Messages []*Message `gorm:"foreignKey:AuthorId"`
	Chats    []*Chat    `gorm:"many2many:chat_members;"`
}

type Message struct {
	base.Base
	Body     string         `json:"body"`
	AuthorId string         `json:"authorId"`
	ChatId   string         `json:"chatId"`
	Seen     pq.StringArray `gorm:"type:text[]" json:"seen"`
	Author   *User          `json:"author"`
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
