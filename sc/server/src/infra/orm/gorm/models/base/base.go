package base

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4();"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	db.Statement.SetColumn("ID", uuid)
	return nil
}
