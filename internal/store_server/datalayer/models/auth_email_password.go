package models

type AuthEmailPassword struct {
	BaseModel
	Email     string  `gorm:"type:varchar(255);unique;not null"`
	Password  string  `gorm:"type:varchar(255);not null"`
	ProfileID uint    `gorm:"type:int;not null"`
	Profile   Profile `gorm:"foreignKey:ProfileID"`
}
