package message

import (
	"src/infra/orm/gorm/models/base"
)

type Message struct {
	base.Base
	Body     string `json:"body"`
	AuthorId string `json:"authorId"`
	Seen     bool   `json:"seen"`
}
