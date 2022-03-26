package post

type Post struct {
	Id, Title, description string
}

func NewPost(id, title, description string) *Post {
	return &Post{id, title, description}
}

func (p *Post) Description() string {
	return p.description
}
