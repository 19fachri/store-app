package models

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ExternalID string `gorm:"unique" json:"external_id"`
}

func (bm *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if strings.EqualFold(bm.ExternalID, "") {
		bm.ExternalID = uuid.New().String()
	}

	return nil
}
