package base

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID      `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.NewV4()
	db.Statement.SetColumn("ID", uuid)
	return db.Error
}
