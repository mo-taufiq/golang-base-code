package model

import (
	"time"

	"gorm.io/gorm"
)

type ModelFieldsDefault struct {
	CreatedAt *time.Time      `json:"created_at" gorm:"<-:create" deepcopier:"field:created_at"`
	UpdatedAt *time.Time      `json:"updated_at" gorm:"<-:update" deepcopier:"field:updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index" deepcopier:"field:deleted_at"`
}
