package chat

import (
	"src/infra/orm/gorm/models/base"
	"src/infra/orm/gorm/models/user"
)

type Chat struct {
	base.Base
	Members []*user.User `gorm:"many2many:chat_members;"`
}
