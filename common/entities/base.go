package entities

import (
	"github.com/jinzhu/gorm"
	. "github.com/satori/go.uuid"
	"time"
)

type Base struct {
	Id        UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := NewV4()
	return scope.SetColumn("id", uuid)
}
