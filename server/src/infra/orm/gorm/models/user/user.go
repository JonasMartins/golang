package user

import (
	base "src/infra/orm/gorm/models/base"
)

type User struct {
	base.Base
	Name     string `json:"name"`
	Email    string `gorm:"index:unique"`
	Password string `json:"_"`
}
