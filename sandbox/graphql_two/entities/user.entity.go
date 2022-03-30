package entities

import "time"

//import "github.com/google/uuid"

type User struct {
	ID        string    `json:id`
	Name      string    `json:name`
	Email     string    `json:email`
	password  string    `json:password`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}

func (u *User) Password() string {
	return u.password
}
