package actions

import (
	"github.com/19fachri/store-app/internal/store_server/datalayer/models"
	"gorm.io/gorm"
)

type ProfileActionInterface interface {
	SaveProfile(data *models.Profile) error
	CountProfileByEMail(email string) (int64, error)
	CountProfileByUsername(username string) (int64, error)
}

type ProfileAction struct {
	db *gorm.DB
}

func NewProfileAction(db *gorm.DB) ProfileActionInterface {
	return &ProfileAction{
		db: db,
	}
}

func (a *ProfileAction) SaveProfile(data *models.Profile) error {
	return a.db.Save(&data).Error
}

func (a *ProfileAction) CountProfileByEMail(email string) (int64, error) {
	var count int64
	err := a.db.Model(&models.Profile{}).Where("email = ?", email).Count(&count).Error
	return count, err
}

func (a *ProfileAction) CountProfileByUsername(username string) (int64, error) {
	var count int64
	err := a.db.Model(&models.Profile{}).Where("username = ?", username).Count(&count).Error
	return count, err
}
