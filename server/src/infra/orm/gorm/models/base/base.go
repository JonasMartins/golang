package base

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	db.Statement.SetColumn("Id", uuid)
	return nil
}
