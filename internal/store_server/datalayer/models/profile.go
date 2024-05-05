package models

type Profile struct {
	BaseModel
	Name     string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
}
