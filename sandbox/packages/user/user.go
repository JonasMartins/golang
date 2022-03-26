package user

import (
	postEntity "packages/post"
)

type User struct {
	Id, Name, Email string
	Posts           []postEntity.Post
}

func NewUser(id, name, email string) *User {
	posts := []postEntity.Post{}
	return &User{id, name, email, posts}
}

func (u *User) AssignPostToUser(p *postEntity.Post) *User {
	u.Posts = append(u.Posts, *p)
	return u
}
