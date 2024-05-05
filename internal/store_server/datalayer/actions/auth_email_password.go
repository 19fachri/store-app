package actions

import (
	"github.com/19fachri/store-app/internal/store_server/datalayer/models"
	"gorm.io/gorm"
)

type AuthEmailPasswordActionInterface interface {
	SaveAuthEmailPassword(data *models.AuthEmailPassword) error
	CountByEmail(email string) (int64, error)
	GetByEmail(email string) (*models.AuthEmailPassword, error)
}

type AuthEmailPasswordAction struct {
	db *gorm.DB
}

func NewAuthEmailPasswordAction(db *gorm.DB) AuthEmailPasswordActionInterface {
	return &AuthEmailPasswordAction{
		db: db,
	}
}

func (a *AuthEmailPasswordAction) SaveAuthEmailPassword(data *models.AuthEmailPassword) error {
	return a.db.Save(&data).Error
}

func (a *AuthEmailPasswordAction) CountByEmail(email string) (int64, error) {
	var count int64
	err := a.db.Model(&models.AuthEmailPassword{}).Where("email = ?", email).Count(&count).Error
	return count, err
}

func (a *AuthEmailPasswordAction) GetByEmail(email string) (*models.AuthEmailPassword, error) {
	var data models.AuthEmailPassword
	err := a.db.Where("email = ?", email).Preload("Profile").First(&data).Error
	return &data, err
}
