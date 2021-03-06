// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"src/infra/orm/gorm/models"
)

type AuthResponse struct {
	Token  string   `json:"token"`
	Errors []*Error `json:"errors"`
}

type CreateAction struct {
	Errors  []*Error `json:"errors"`
	Created bool     `json:"created"`
}

type CreateChatInput struct {
	Messages []string `json:"messages"`
	Members  []string `json:"members"`
}

type CreateMessageInput struct {
	Body     string   `json:"body"`
	AuthorID string   `json:"authorId"`
	ChatID   string   `json:"chatId"`
	Seen     []string `json:"seen"`
}

type CreteAction struct {
	Errors  []*Error `json:"errors"`
	Created bool     `json:"created"`
}

type DeleteAction struct {
	Message string   `json:"message"`
	Status  string   `json:"status"`
	Errors  []*Error `json:"errors"`
}

type Error struct {
	Method  string `json:"method"`
	Message string `json:"message"`
	Field   string `json:"field"`
	Code    int    `json:"code"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MessagesResponse struct {
	Errors   []*Error          `json:"errors"`
	Messages []*models.Message `json:"messages"`
}

type RegisterUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	User   *models.User `json:"user"`
	Errors []*Error     `json:"errors"`
}

type UsersResponse struct {
	Users  []*models.User `json:"users"`
	Errors []*Error       `json:"errors"`
}
